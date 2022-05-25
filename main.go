package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rhmdnd/compserv-poc/compserv"
	pb "github.com/rhmdnd/compserv-poc/compserv"
)

func main() {
	// Get a connection to the database so we can pass it in as a
	// dependency to the service implementations.
	c := compserv.ParseConfig()
	connStr := compserv.GetDatabaseConnectionString(c)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Established database connection. %v", db)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAssessmentResultServer(grpcServer, compserv.NewServer(db))
	log.Printf("gRPC server listening on %s", lis.Addr())
	grpcServer.Serve(lis)
}
