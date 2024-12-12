package server

import (
	"context"
	"log"
	"net"

	pb "github.com/a-korkin/notes_ms/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type NoteServer struct {
	pb.UnimplementedNoteServiceServer
}

func (s *NoteServer) AddNote(
	ctx context.Context, noteIn *pb.NoteIn) (*pb.Note, error) {
	log.Printf("add note method called")
	note := pb.Note{
		Id:   1,
		Text: "Note1",
	}
	return &note, nil
}

func (s *NoteServer) GetNote(
	ctx context.Context, noteId *pb.NoteId) (*pb.Note, error) {
	log.Printf("get node method called")
	note := pb.Note{
		Id:   1,
		Text: "Note1",
	}
	return &note, nil
}

func NewServer() *NoteServer {
	return &NoteServer{}
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