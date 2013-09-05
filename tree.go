package RiverFS

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	//"github.com/goraft/raft"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"syscall"
)

type RootDir struct {
	root Dir
}

func (nf TargetDir) Root() (fs.Node, fuse.Error) {
	return root, nil
}

type Node struct {
	Name string
}

func (n Node) Attr() fuse.Attr {
	s, err := os.Stat(n.Path)
	if err != nil {
		log.Print(err)
		return fuse.Attr{}
	}

	return fuse.Attr{Size: uint64(s.Size()), Mtime: s.ModTime(), Mode: s.Mode()}
}

type Dir struct {
	Node
	rwmut    sync.RWMutex
	Children map[string]Node
}

func (d Dir) Lookup(name string, intr fs.Intr) (fs fs.Node, err fuse.Error) {
	d.rwmut.RLock()
	node, ok := d.Children[name]
	d.rwmut.RUnlock()
	if !ok {
		return nil, fuse.ENOENT
	}

	return
}

func (d Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	var out []fuse.Dirent

	d.rwmut.RLock()
	for name, node := range d.Children {
		de := fuse.Dirent{Name: name}
		switch node.(type) {
		case Dir:
			de.Type = fuse.DT_Dir
		case File:
			de.Type = fuse.DT_Dir
		}
		out = append(out, de)
	}
	d.rwmut.RUnlock()

	return out, nil
}

type File struct {
	Node
	Host string
	Path string
}

func (f File) ReadAll(intr fs.Intr) ([]byte, fuse.Error) {
	contents, err := ioutil.ReadFile(Path)
	return contents, err
}
