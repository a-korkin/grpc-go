package server

import (
	"context"
	"log"
	"net"

	pb "github.com/a-korkin/notes_ms/common"
	"github.com/a-korkin/notes_ms/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type NoteServer struct {
	pb.UnimplementedNoteServiceServer
	dbConnect *data.DbConnect
}

func NewServer(dbconn *data.DbConnect) *NoteServer {
	return &NoteServer{dbConnect: dbconn}
}

func (s *NoteServer) AddNote(
	ctx context.Context, noteIn *pb.NoteIn) (*pb.Note, error) {
	log.Printf("add note method called")
	note, err := s.dbConnect.AddNote(noteIn)
	if err != nil {
		log.Printf("failed to create note: %s", err)
		return nil, err
	}
	return note, err
}

func (s *NoteServer) GetNote(
	ctx context.Context, noteId *pb.NoteId) (*pb.Note, error) {
	log.Printf("get node method called")
	note, err := s.dbConnect.GetNote(noteId.Id)
	if err != nil {
		log.Printf("failed to get note with id: %d", noteId.Id)
		return nil, err
	}
	return note, nil
}

func (s *NoteServer) Run(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to create connection: %s", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNoteServiceServer(grpcServer, s)

	reflection.Register(grpcServer)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
