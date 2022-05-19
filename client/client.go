package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/rhmdnd/compserv-poc/compserv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost", "Hostname or IP address of the server.")
	serverPort = flag.String("port", "50051", "Port the server is listening on.")
	subject    = flag.String("subject", "", "Result subject (e.g., cluster or node).")
	control    = flag.String("control", "", "Compliance control (e.g., ac-1.1.2).")
	rule       = flag.String("rule", "", "OpenSCAP rule name (e.g., content_rule_usbguard).")
	outcome    = flag.String("outcome", "", "Rule outcome (e.g., PASS, FAIL).")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*serverAddr+":"+*serverPort, opts...)
	if err != nil {
		log.Fatalf("Unable to establish connection to %s:%s: %v", *serverAddr, *serverPort, err)
	}
	defer conn.Close()
	client := pb.NewAssessmentResultClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := &pb.Result{Subject: *subject, Control: *control, Rule: *rule, Outcome: *outcome}
	client.AddAssessmentResult(ctx, result)
}
