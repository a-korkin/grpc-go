package data_test

import (
	"fmt"
	"github.com/a-korkin/notes_ms/data"
	"log"
	"os"
	"testing"
)

var DbConnect *data.DbConnect

func connectToDb(dbName *string) {
	conn := fmt.Sprintf(
		"user=postgres password=admin dbname=%s sslmode=disable", *dbName)
	DbConnect = data.NewDbConnect(conn)
}

func setup(dbName *string) {
	log.Printf("setup...")
	DbConnect.Db.Exec(fmt.Sprintf("create database %s;", *dbName))
	connectToDb(dbName)
	DbConnect.Db.Exec(`
create table if not exists notes
(
	id serial primary key,
	text varchar(255)
);
insert into notes(text)
values('first note'), ('second note'), ('third note');
`)
}

func shutdown() {
	log.Printf("shutdown...")
	DbConnect.Db.Exec(fmt.Sprintf("drop table if exists notes;"))
}

func TestMain(m *testing.M) {
	dbName := "postgres"
	connectToDb(&dbName)

	dbName = "notes_db_test"
	setup(&dbName)
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestGetNoteSucces(t *testing.T) {
	note, err := DbConnect.GetNote(1)
	if note.Text != "first note" {
		t.Errorf("wrong note getted")
	}
	if err != nil {
		t.Errorf("error of getting note: %s", err)
	}
}
