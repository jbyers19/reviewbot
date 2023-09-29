package main

import (
	"fmt"
	"log"
	"net"
	"os"

	rb "github.com/jbyers19/reviewbot/chatbot/internal/app/reviewbot"
	"github.com/jbyers19/reviewbot/chatbot/internal/app/reviewbot/pb"
	"google.golang.org/grpc"
)

// TemplatesDB and CustomerDB holds templates and customer data.
// In a real application, this data should be stored in a database.
// Also, this is not thread-safe, but it's fine for this practice project.
var TemplatesDB *rb.Templates
var CustomerDB *rb.Customers

func init() {
	log.Println("initializing Templates Database...")
	TemplatesDB = &rb.Templates{TemplatesMap: make(map[string]*pb.MessageTemplate)}

	log.Println("initializing Customer Database...")
	CustomerDB = &rb.Customers{CustomersMap: make(map[int64]*rb.Customer)}
}

func main() {
	// Start the gRPC server and the Telegram bot.
	go StartGRPCServer()
	rb.StartTelegramBot(TemplatesDB, CustomerDB)
}

// StartGRPCServer starts the gRPC server and registers the Templates and Message services.
func StartGRPCServer() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("starting gRPC server on port %s", port)
	srv := grpc.NewServer()
	pb.RegisterTemplatesServer(srv, TemplatesDB)

	msgsrv := &rb.Message{}
	pb.RegisterMessageServer(srv, msgsrv)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
