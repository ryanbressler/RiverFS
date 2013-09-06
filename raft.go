package RiverFS

import (
	"github.com/goraft/raft"
	"log"
	"net/http"
)

type StateMachine struct{}

func (s *StateMachine) Save() ([]byte, error) {
	return nil, nil
}

func (s *StateMachine) Recovery(bs []byte) error {
	return nil
}

func StartRaftServer(me string, path string, lead bool, peers []string) {
	log.Println("start raft")

	transporter := raft.NewHTTPTransporter("raft")

	raftserver, err := raft.NewServer(me, path, transporter, &StateMachine{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	httpserver := &http.Server{
		Handler: mux,
		Addr:    me,
	}

	transporter.Install(raftserver, mux)

	//raftserver.Start()
	log.Println("listen and serve")
	httpserver.ListenAndServe()

}
