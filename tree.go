package RiverFS

import (
	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	//"github.com/goraft/raft"
	"os"
	"sync"
)

type RootDir struct {
	Dir *Dir
}

func NewRootDir() (r *RootDir) {
	r = &RootDir{Dir: NewDir()}
	return
}

func (nf *RootDir) Root() (fs.Node, fuse.Error) {
	return nf.Dir, nil
}

type Node struct {
	Name string
	Mode os.FileMode
}

func (n *Node) Attr() fuse.Attr {
	return fuse.Attr{Mode: n.Mode}
}

type Dir struct {
	Node
	rwmut    sync.RWMutex
	Children map[string]fs.Node
}

func NewDir() (d *Dir) {
	d = &Dir{}
	d.Node = Node{}
	d.Mode = os.ModeDir | 0555
	d.Children = make(map[string]fs.Node)
	return
}

func (d *Dir) Mkdir(req *fuse.MkdirRequest, intr fs.Intr) (n fs.Node, ferr fuse.Error) {
	ndir := NewDir()
	ndir.Name = req.Name

	//TODO: consensous
	//TODO: Check if exists and error

	d.rwmut.Lock()
	d.Children[req.Name] = ndir
	d.rwmut.Unlock()

	return ndir, nil

}

func (d *Dir) Lookup(name string, intr fs.Intr) (fs fs.Node, err fuse.Error) {
	d.rwmut.RLock()
	fs, ok := d.Children[name]
	d.rwmut.RUnlock()
	if !ok {
		return nil, fuse.ENOENT
	}

	return
}

func (d *Dir) ReadDir(intr fs.Intr) ([]fuse.Dirent, fuse.Error) {
	var out []fuse.Dirent

	d.rwmut.RLock()
	for name, node := range d.Children {
		de := fuse.Dirent{Name: name}
		switch node.(type) {
		case *Dir:
			de.Type = fuse.DT_Dir
		case *File:
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

func (f *File) ReadAll(intr fs.Intr) (out []byte, ferr fuse.Error) {
	//Place holder code
	out = make([]byte, 0)
	return out, ferr
}
