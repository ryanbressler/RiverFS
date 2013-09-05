package RiverFS

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	//"github.com/goraft/raft"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

func ServeDir(mountpoint string, datadir string, me string, peerlist []string) {

	c, err := fuse.Mount(mountpoint)
	if err != nil {
		log.Fatal(err)
	}

	fs.Serve(c, TargetDir{datadir})

}
