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

func (dbConnect *DbConnect) GetNote(id int32) (*common.Note, error) {
	row, err := dbConnect.Db.Query("select text from notes where id = $1", id)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		note := common.Note{Id: id}
		row.Scan(&note.Text)
		return &note, nil
	}
	return nil, nil
}

func (dbConnect *DbConnect) AddNote(in *common.NoteIn) (*common.Note, error) {
	row, err := dbConnect.Db.Query(`
insert into notes(text)
values($1)
returning id`, in.Text)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		note := common.Note{}
		row.Scan(&note.Id, &note.Text)
		return &note, nil
	}
	return nil, nil
}
