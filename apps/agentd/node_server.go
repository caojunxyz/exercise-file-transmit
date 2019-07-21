package main

import (
	"context"
	"encoding/json"
	"github.com/caojunxyz/superdeliver/deliver"
	"log"
)

type NodeServer struct {
	dispatcher *BlockDispatcher
}

func (ns *NodeServer) InitSession(ctx context.Context, arg *deliver.Session) (*deliver.Null, error) {
	data, _ := json.Marshal(arg)
	log.Printf("%s\n", string(data))
	ns.dispatcher = NewBlockDispatcher(arg)
	return &deliver.Null{}, nil
}

func (ns *NodeServer) NotifyBlockFlag(ctx context.Context, arg *deliver.BlockFlag) (*deliver.Null, error) {
	ns.dispatcher.BlockFlagReceived(arg)
	return &deliver.Null{}, nil
}

func (ns *NodeServer) PushBlock(ctx context.Context, arg *deliver.Block) (*deliver.Null, error) {
	ns.dispatcher.BlockReceived(arg)
	return &deliver.Null{}, nil
}

func (ns *NodeServer) PullBlock(ctx context.Context, arg *deliver.BlockFlag) (*deliver.Block, error) {
	return ns.dispatcher.Assembler.ReadBlock(arg.BlockId)
}

func (ns *NodeServer) CheckProgress(ctx context.Context, arg *deliver.Progress) (*deliver.Progress, error) {
	return ns.dispatcher.CheckProgress(arg.SessionId), nil
}

