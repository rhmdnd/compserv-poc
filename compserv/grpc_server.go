package compserv

import (
	context "context"
	"log"

	"gorm.io/gorm"
)

type resultServer struct {
	UnimplementedAssessmentResultServer
	database *gorm.DB
}

func NewServer(db *gorm.DB) *resultServer {
	s := &resultServer{database: db}
	return s
}

func (r *resultServer) AddAssessmentResult(ctx context.Context, result *Result) (*Result, error) {
	log.Printf("Received result: %s", result)

	// TODO(rhmdnd): Implement storage using the provided database interface
	return &Result{}, nil
}
