package main

import (
	"fmt"
	"local-storage-server/internal/server"
)

func main() {

	newServer := server.NewServer()

	err := newServer.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start newServer: %s", err))
	}
}
