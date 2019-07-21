package main

import (
	"context"
	"github.com/caojunxyz/superdeliver/deliver"
	"github.com/caojunxyz/superdeliver/fileblock"
	"google.golang.org/grpc"
	"log"
	"sync"
)

// when a block is pushed to node, then the node does these following jobs by order :
// 1. assemble block
// 2. broadcast this block flag (which means it has this block) to some other nodes according to its role:
//    3.1. if it's a leader, broadcast to other group leaders and one of the normal node in the same group
//    3.2. if it's a normal node, broadcast to other nodes in the same group except the leader
// 3. if a node receives a block flag message, check whether it has this block first, if not, pull the block from the node
//    who has it.
// 4. if a normal node receives all blocks and completes the assembling, notifies its group leader. And if the leader
//    receives completes, notifies the master (manager).
type BlockDispatcher struct {
	sync.RWMutex
	Session       *deliver.Session
	Assembler     *fileblock.AssembleController
	pickNodeIdx   int
	curNode       *deliver.Node
	chanAssembled chan *deliver.BlockAssembled // finish assembling, wait for broadcast
	nodesClient   map[int64]deliver.DeliverClient
}

func NewBlockDispatcher(session *deliver.Session) *BlockDispatcher {
	ch := make(chan *deliver.BlockAssembled, 50)
	dp := &BlockDispatcher{
		Session:       session,
		Assembler:     fileblock.NewAssembleController(session.Artifact, ch),
		chanAssembled: ch,
		nodesClient:   make(map[int64]deliver.DeliverClient),
	}

	for _, n := range dp.Session.Group.Nodes {
		if n.Id == dp.Session.NodeId {
			dp.curNode = n
			break
		}
	}

	if dp.curNode == nil {
		log.Panic("should not be here:", dp.Session)
	}

	go dp.broadcast()
	return dp
}

func (dp *BlockDispatcher) getNodeClient(node *deliver.Node) deliver.DeliverClient {
	dp.RLock()
	client, ok := dp.nodesClient[node.Id]
	dp.RUnlock()

	if !ok {
		dp.Lock()
		defer dp.Unlock()
		conn, err := grpc.Dial(node.Addr, grpc.WithInsecure())
		if err != nil {
			log.Panic("did not connect: %v", err)
		}

		client = deliver.NewDeliverClient(conn)
		dp.nodesClient[node.Id] = client
	}
	return client
}

func (dp *BlockDispatcher) pull(bf *deliver.BlockFlag) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	client := dp.getNodeClient(bf.Node)
	block, err := client.PullBlock(context.Background(), bf)
	if err != nil {
		log.Panic(err)
		return
	}

	if err := fileblock.CheckBlock(block); err != nil {
		log.Panic(err)
	}

	block.PulledFrom = []int64{bf.Node.Id}
	log.Printf("Pulled block#%d from node-%d\n", block.Id, bf.Node.Id)
	dp.Assembler.SendToAssemble(block)
}

func (dp *BlockDispatcher) BlockFlagReceived(bf *deliver.BlockFlag) {
	if !dp.Assembler.IsBlockAssembled(bf.BlockId) {
		go dp.pull(bf)
	}
}

func (dp *BlockDispatcher) getBroadcastReceivers(ba *deliver.BlockAssembled) []*deliver.Node {
	receivers := []*deliver.Node{}
	switch dp.curNode.Type {
	case deliver.NodeType_LEADER:
		// all other leaders except the "pulled from" leader
		for _, v := range dp.Session.Leaders {
			if v.Id == dp.curNode.Id {
				continue
			}
			if len(ba.PulledFrom) > 0 && ba.PulledFrom[0] == v.Id {
				continue
			}
			receivers = append(receivers, v)
		}

		list := []*deliver.Node{}
		for _, n := range dp.Session.Group.Nodes {
			if dp.curNode.Id == n.Id {
				continue
			}
			list = append(list, n)
		}
		if len(list) > 0 {
			idx := int(ba.Id % int64(len(list)))
			receivers = append(receivers, list[idx])
		}
	case deliver.NodeType_NORMAL:
		for _, v := range dp.Session.Group.Nodes {
			if dp.curNode.Id == v.Id || v.Type == deliver.NodeType_LEADER {
				continue
			}
			if len(ba.PulledFrom) > 0 && ba.PulledFrom[0] == v.Id {
				continue
			}
			receivers = append(receivers, v)
		}
	}
	return receivers
}

func (dp *BlockDispatcher) broadcast() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("broadcast...")
	for {
		select {
		case bm := <-dp.chanAssembled:
			bf := &deliver.BlockFlag{BlockId: bm.Id, Node: dp.curNode}
			receivers := dp.getBroadcastReceivers(bm)
			for _, node := range receivers {
				log.Printf("Broadcast block#%d flag from node-%d to node-%d\n", bf.BlockId, bf.Node.Id, node.Id)
				client := dp.getNodeClient(node)
				_, err := client.NotifyBlockFlag(context.Background(), bf)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func (dp *BlockDispatcher) BlockReceived(block *deliver.Block) error {
	dp.Assembler.SendToAssemble(block)
	return nil
}

func (dp *BlockDispatcher) CheckProgress(sessionId int64) *deliver.Progress {
	result := &deliver.Progress{SessionId:sessionId, AssembledBlocks: make(map[int64]int64)}
	switch dp.curNode.Type {
	case deliver.NodeType_LEADER:
		result.AssembledBlocks[dp.curNode.Id] = dp.Assembler.Progress()
		for _, v := range dp.Session.Group.Nodes {
			if v.Id != dp.curNode.Id {
				client := dp.getNodeClient(v)
				progress, err := client.CheckProgress(context.Background(), &deliver.Progress{SessionId: sessionId})
				if err != nil {
					log.Println(err)
					continue
				}
				result.AssembledBlocks[v.Id] = progress.AssembledBlocks[v.Id]
			}
		}
	case deliver.NodeType_NORMAL:
		result.AssembledBlocks[dp.curNode.Id] = dp.Assembler.Progress()
	}
	return result
}
