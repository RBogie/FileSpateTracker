package main

import (
	"fmt"

	"github.com/robbiedobbie/FileSpateTracker/server"
)

func main() {
	fmt.Println("Hello all!")
	s := server.NewServer()
	s.Listen()
}
