package main

import (
	"context"
	"fmt"
	"github.com/caojunxyz/superdeliver/deliver"
	"github.com/caojunxyz/superdeliver/fileblock"
	"github.com/caojunxyz/superdeliver/utils"
	"google.golang.org/grpc"
	"log"
	"path/filepath"
	"time"
)

const (
	DefGroupNum = 10
)

// node id starts from 0
// group id starts from 0
type DeliverManager struct {
	SessionId     int64
	Groups        []*deliver.Group
	LeadersClient map[int64]deliver.DeliverClient
	Splitter      *fileblock.SplitController
}

func (m *DeliverManager) GetLeaders() []*deliver.Node {
	leaders := []*deliver.Node{}
	for _, group := range m.Groups {
		for _, v := range group.Nodes {
			if group.LeaderId == v.Id {
				leaders = append(leaders, v)
				break
			}
		}
	}
	return leaders
}

func (m *DeliverManager) GetLeaderGroup(leaderId int64) *deliver.Group {
	for _, g := range m.Groups {
		if g.LeaderId == leaderId {
			return g
		}
	}
	return nil
}

func (m *DeliverManager) InitGroups(servers []string) {
	if len(servers) == 0 {
		return
	}

	m.LeadersClient = make(map[int64]deliver.DeliverClient)
	groupNum := DefGroupNum
	groupSize := len(servers) / DefGroupNum
	if groupSize <= 0 {
		groupSize = 1
		groupNum = len(servers)
	} else {
		groupNum += (len(servers) % DefGroupNum)
	}

	srvIdx := 0
	maxNodeId := int64(len(servers))
	for i := 0; i < groupNum; i++ {
		group := &deliver.Group{Id: int64(i)}
		for {
			srv := servers[srvIdx]
			node := &deliver.Node{Id: int64(srvIdx), Type: deliver.NodeType_NORMAL, Addr: srv}
			if len(group.Nodes) == 0 {
				node.Type = deliver.NodeType_LEADER
				group.LeaderId = node.Id
			}
			group.Nodes = append(group.Nodes, node)
			srvIdx++
			if len(group.Nodes) == groupSize || node.Id == maxNodeId || srvIdx >= len(servers) {
				break
			}
		}
		m.Groups = append(m.Groups, group)
	}
	fmt.Println("------------------ groups ------------------")
	for _, g := range m.Groups {
		fmt.Printf("group#%d:\nLeader: %d (%s)\n", g.Id, g.LeaderId, g.Nodes[0].Addr)
		for _, v := range g.Nodes {
			if v.Type == deliver.NodeType_NORMAL {
				fmt.Printf("Node: %d (%s)\n", v.Id, v.Addr)
			}
		}
	}
}

func (m *DeliverManager) InitSessions(filename string) {
	leaders := m.GetLeaders()
	m.Splitter = fileblock.NewSplitController(filename, len(leaders))
	if err := m.Splitter.Init(); err != nil {
		log.Panic(err)
	}

	for _, leader := range leaders {
		conn, err := grpc.Dial(leader.Addr, grpc.WithInsecure())
		if err != nil {
			log.Panic("did not connect: %v", err)
		}

		client := deliver.NewDeliverClient(conn)
		group := m.GetLeaderGroup(leader.Id)
		artifact := &deliver.Artifact{
			Filename:  filepath.Base(m.Splitter.Artifact.Filename),
			Size:      m.Splitter.Artifact.Size,
			Sha1:      m.Splitter.Artifact.Sha1,
			BlockNum:  m.Splitter.Artifact.BlockNum,
			BlockSize: m.Splitter.Artifact.BlockSize,
		}
		session := &deliver.Session{SessionId: m.SessionId, Artifact: artifact, NodeId: leader.Id, Group: group, Leaders: leaders}
		_, err = client.InitSession(context.Background(), session)
		if err != nil {
			log.Panic("init session: %v", err)
		}
		m.LeadersClient[leader.Id] = client
	}
}

func (m *DeliverManager) Start() {
	if err := m.Splitter.Start(); err != nil {
		log.Panic(err)
	}

	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case block := <-m.Splitter.Blocks:
			// log.Println(utils.BlockFormatString(block))
			group := m.Groups[int(block.Id%int64(len(m.Groups)))]
			client := m.LeadersClient[group.LeaderId]
			// log.Printf("Push block#%d to node-%d\n", block.Id, group.LeaderId)
			_, err := client.PushBlock(context.Background(), block)
			if err != nil {
				log.Panic(err)
			}

			if block.Id == m.Splitter.Artifact.BlockNum-1 {
				log.Println("Done!")
			}

		case <-ticker.C:
			for _, client := range m.LeadersClient {
				progress, err := client.CheckProgress(context.Background(), &deliver.Progress{SessionId: m.SessionId})
				if err == nil {
					log.Println(utils.ProgressFormatString(progress, m.Splitter.Artifact.BlockNum))
				}
			}
		}
	}
}
