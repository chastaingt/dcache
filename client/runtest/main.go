package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chastaingt/dcache/client"
)

func main() {
	// var (
	// 	cfg           = raft.DefaultConfig()
	// 	fsm           = &raft.MockFSM{}
	// 	logStore      = raft.NewInmemStore()
	// 	stableStore   = raft.NewInmemStore()
	// 	snapShotStore = raft.NewInmemSnapshotStore()
	// 	timeout       = time.Second * 5
	// )

	// cfg.LocalID = "toto"

	// // ips, err := net.LookupIP("localhost")
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // if len(ips) == 0 {
	// // 	log.Fatalf("localhost did not resolve any ips")
	// // }
	// // addr := &net.TCPAddr{IP: ips[0], Port: 4000}

	// tr, err := raft.NewTCPTransport("127.0.0.1:4000", nil, 10, timeout, os.Stdout)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// server := raft.Server{
	// 	Suffrage: raft.Voter,
	// 	ID:       raft.ServerID(cfg.LocalID),
	// 	Address:  raft.ServerAddress("127.0.0.1:4000"),
	// }

	// serverConfig := raft.Configuration{
	// 	Servers: []raft.Server{server},
	// }
	// r, err := raft.NewRaft(cfg, fsm, stableStore, logStore, snapShotStore, tr)
	// if err != nil {
	// 	log.Fatal("Failed to create new raft", err)
	// }

	// if err := r.BootstrapCluster(serverConfig).Error(); err != nil {
	// 	log.Fatal("Failed to Boostrap cluster", err)
	// }
	// fmt.Printf("%+v\n", r)

	// select {}
	SendStuff()
}

func SendStuff() {
	client, err := client.New(":3000", client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 100; i++ {
		var (
			key   = []byte(fmt.Sprintf("key_%d", i))
			value = []byte(fmt.Sprintf("val_%d", i))
		)
		if err := client.Set(context.Background(), key, value, 0); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
	client.Close()
}
