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
	Root Dir
}

func NewRootDir() (r *RootDir) {
	r = &RootDir{
		Root: Dir{
			Node{Name: "",
				Attr: fuse.Attr{}},
			Children: make(map[string]Node)}}
}

func (nf TargetDir) Root() (fs.Node, fuse.Error) {
	return Root, nil
}

type Node struct {
	Name string
	Attr fuse.Attr
}

func (n Node) Attr() fuse.Attr {

	return n.Attr
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

type Chunk struct {
	Host string
	Path string
}

type File struct {
	Node
	Chunks []Chunk
}

func (f File) ReadAll(intr fs.Intr) ([]byte, fuse.Error) {
	//Place holder code ...not effichent and only supports single machine.
	out = ""
	ferr = nil
	for _, c := range f.Chunks {
		contents, err := ioutil.ReadFile(c.Path)
		if err {
			ferr = fuse.EIO
		}
		out += contents
	}
	return out, ferr
}
