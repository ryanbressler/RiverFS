package RiverFS

import (
	"github.com/goraft/raft"
	"log"
)

type StateMachine struct{}

func (s *StateMachine) Save() ([]byte, error) {
	return nil, nil
}

func (s *StateMachine) Recovery(bs []byte) error {
	return nil
}

func StartRaftServer(me string, path string, lead bool, peers []string) {

	transporter := raft.NewHTTPTransporter("raft")

	server, err := raft.NewServer(me, path, transporter, &StateMachine{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	server.Start()

}
