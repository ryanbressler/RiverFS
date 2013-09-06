package main

import (
	"encoding/csv"
	"flag"
	"github.com/ryanbressler/RiverFS"
	"io"
	"log"
	"os"
)

func main() {
	var mountpoint string
	flag.StringVar(&mountpoint, "mountpoint", "river", "Where to mount the fuse dir.")

	var raftlog string
	flag.StringVar(&raftlog, "raftlog", "raftlog", "The path to the local raft log.")

	var me string
	flag.StringVar(&me, "me", "http://localhost", "Address for this node.")

	var peerlist string
	flag.StringVar(&peerlist, "peers", "", "Text file to load peers from.")

	var lead bool
	flag.BoolVar(&lead, "lead", false, "Start as leader.")

	peers := make([]string, 10)
	if peerlist != "" {

		peerfile, err := os.Open(peerlist)
		if err != nil {
			log.Fatal(err)
		}
		tsv := csv.NewReader(peerfile)
		tsv.Comma = '\t'
		for {
			url, err := tsv.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			peers = append(peers, url[0])

		}
		peerfile.Close()

	}

	flag.Parse()

	//RiverFS.StartRaftServer(me, raftlog, lead, peers)
	RiverFS.FuseMount(mountpoint)

}
