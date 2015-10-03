package server

import (
	"fmt"
	"net/http"

	"github.com/robbiedobbie/FileSpateTracker/protogen"

	"github.com/golang/protobuf/proto"
)

type server struct {
}

//NewServer will create a new server which can be used to listen to incoming connections
func NewServer() *server {
	s := new(server)
	return s
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rawMessage := make([]byte, r.ContentLength)
	r.Body.Read(rawMessage)
	message := new(protogen.Request)
	proto.Unmarshal(rawMessage, message)
	fmt.Fprint(w, "Nobody should read this.")
}

func youHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You should read this.")
}

//Will make the server listen for new connections
func (s server) Listen() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/you", youHandler)
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Printf("main(): %s\n", err)
	}
}
