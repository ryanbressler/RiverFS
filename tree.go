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

func (n *Node) Setattr(req *fuse.SetattrRequest, resp *fuse.SetattrResponse, intr fs.Intr) fuse.Error {
	n.Mode = req.Mode
	return nil
}

type Dir struct {
	Node
	rwmut    sync.RWMutex
	Children map[string]fs.Node
}

func NewDir() (d *Dir) {
	d = &Dir{}
	d.Node = Node{}
	d.Mode = os.ModeDir
	d.Children = make(map[string]fs.Node)
	return
}

func (d *Dir) Add(name string, n fs.Node) {
	//TODO: consensous
	//TODO: Check if exists and error
	d.rwmut.Lock()
	d.Children[name] = n
	d.rwmut.Unlock()

}

func (d *Dir) Rm(name string) {
	//TODO: consensous
	//TODO: Check if exists and error
	//TODO: clean up disk space
	d.rwmut.Lock()
	delete(d.Children, name)
	d.rwmut.Unlock()

}

func (d *Dir) Mkdir(req *fuse.MkdirRequest, intr fs.Intr) (n fs.Node, ferr fuse.Error) {
	ndir := NewDir()
	ndir.Name = req.Name
	ndir.Mode = req.Mode
	d.Add(req.Name, ndir)
	return ndir, nil

}

func (d *Dir) Remove(req *fuse.RemoveRequest, intr fs.Intr) fuse.Error {
	d.Rm(req.Name)
	return nil
}

func (d *Dir) Rename(req *fuse.RenameRequest, newDir fs.Node, intr fs.Intr) fuse.Error {

	d.rwmut.RLock()
	node, ok := d.Children[req.OldName]
	d.rwmut.RUnlock()
	if !ok {
		return fuse.ENOENT
	}
	d.Rm(req.OldName)

	switch node.(type) {
	case *Dir:
		node.(*Dir).Name = req.NewName
	case *File:
		node.(*File).Name = req.NewName
	}

	newDir.(*Dir).Add(req.NewName, node)

	return nil

}

func (d *Dir) Create(req *fuse.CreateRequest, resp *fuse.CreateResponse, intr fs.Intr) (fs.Node, fs.Handle, fuse.Error) {

	n := &File{}
	n.Mode = req.Mode
	n.Name = req.Name
	d.Add(req.Name, n)
	return n, n, nil
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

type File struct {
	Node
}

func (f *File) ReadAll(intr fs.Intr) (out []byte, ferr fuse.Error) {
	//Place holder code
	return []byte("I AM A FILE!!!!11!!!\n"), ferr
}

func (f *File) WriteAll(in []byte, intr fs.Intr) fuse.Error {
	return nil
}
