package server

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"github.com/Robbiedobbie/FileSpateTracker/protogen"

	"github.com/golang/protobuf/proto"
)

type server struct {
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

//NewServer will create a new server which can be used to listen to incoming connections
func NewServer() *server {
	s := new(server)
	return s
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_ = "breakpoint"
	rawMessage := make([]byte, r.ContentLength)
	r.Body.Read(rawMessage)
	message := new(protogen.Request)
	proto.Unmarshal(rawMessage, message)

	pub, errParse := x509.ParsePKIXPublicKey(message.Data.PeerPublicKey)
	dataBytes, _ := proto.Marshal(message.GetData())
	err := rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), 0, dataBytes, message.Signature)
	fmt.Println("verify:", err)

	if errParse != nil {
		exit(errParse)
	}

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
