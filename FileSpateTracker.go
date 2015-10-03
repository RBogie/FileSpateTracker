package main

import (
	"fmt"

	"github.com/Robbiedobbie/FileSpateTracker/server"
)

func main() {
	fmt.Println("Hello all!")
	s := server.NewServer()
	s.Listen()
}
