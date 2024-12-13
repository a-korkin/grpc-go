package main

import (
	"fmt"
	"github.com/a-korkin/notes_ms/config"
	"github.com/a-korkin/notes_ms/data"
	"github.com/a-korkin/notes_ms/server"
)

func main() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_USR"), config.GetEnv("DB_PWD"), config.GetEnv("DB_NAME"))
	dbConnect := data.NewDbConnect(connStr)
	srv := server.NewServer(dbConnect)
	srv.Run(":8080")
}
