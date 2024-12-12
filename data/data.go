package data

import (
	"database/sql"
	"github.com/a-korkin/notes_ms/common"
	_ "github.com/lib/pq"
	"log"
)

type DbConnect struct {
	Db *sql.DB
}

func NewDbConnect(connStr string) *DbConnect {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to create connection to postgres: %s", err)
	}
	return &DbConnect{Db: db}
}

func GetNote(id int32) *common.Note {
	return &common.Note{
		Id:   id,
		Text: "hello",
	}
}
