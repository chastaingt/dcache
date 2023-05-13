package main

import (
	"flag"
	"strings"

	"github.com/chastaingt/dcache/cache"
)

func main() {
	var (
		listenAddr  = flag.String("listenaddr", ":3000", "listen address of the server")
		raftAddr    = flag.String("raftaddr", ":3001", "lis address of raft")
		leaderAddr  = flag.String("leaderaddr", "", "listen address of the leader")
		raftNodes   = flag.String("raftnode", "", "listen address of other raft nodes")
		localRaftId = flag.String("localraftid", "", "id identification for raft node")
	)
	flag.Parse()

	opts := ServerOpts{
		ListenAddr:  *listenAddr,
		IsLeader:    len(*leaderAddr) == 0,
		LeaderAddr:  *leaderAddr,
		RaftAdr:     *raftAddr,
		RaftNodes:   strings.Split(*raftNodes, "/"),
		LocalRaftId: *localRaftId,
	}
	server := NewServer(opts, cache.New())
	server.Start()
}
