package data_test

import (
	"fmt"
	"github.com/a-korkin/notes_ms/data"
	"log"
	"os"
	"testing"
)

func setup(dbConn *data.DbConnect, dbName *string) {
	log.Printf("setup...")
	dbConn.Db.Exec(fmt.Sprintf("create database %s;", *dbName))
}

func shutdown(dbConn *data.DbConnect, dbName *string) {
	log.Printf("shutdown...")
	dbConn.Db.Exec(fmt.Sprintf("drop database %s;", *dbName))
}

func TestMain(m *testing.M) {
	conn := "user=postgres password=admin sslmode=disable"
	log.Printf("conn: %s", conn)
	dbName := "notes_db_test"
	dbConnect := data.NewDbConnect(conn)
	setup(dbConnect, &dbName)
	code := m.Run()
	shutdown(dbConnect, &dbName)
	os.Exit(code)
}

// func (dbConnect *DbConnect) GetNote(id int32) (*common.Note, error) {
// 	row, err := dbConnect.Db.Query("select text from notes where id = $1", id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if row.Next() {
// 		note := common.Note{Id: id}
// 		row.Scan(&note.Text)
// 		return &note, nil
// 	}
// 	return nil, nil
// }
