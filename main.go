package main

import (
	"github.com/a-korkin/notes_ms/server"
)

func main() {
	srv := server.NewServer()
	srv.Run(":8080")
}
