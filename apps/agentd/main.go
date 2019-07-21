package main

import (
	"flag"
	"github.com/caojunxyz/superdeliver/deliver"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "port", "10000", "listen port")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ldate)

	log.Println("agentd listen at:", port)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	deliver.RegisterDeliverServer(s, &NodeServer{})
	if err := s.Serve(lis); err != nil {
		log.Panic(err)
	}
}

