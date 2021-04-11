package main

import (
	"github.com/datagentleman/cupcake/tcp"
)

func main() {
	err := tcp.StartServer("127.0.0.1", 2001)
	if err != nil {
		panic(err)
	}
}
