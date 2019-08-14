package main

import (
	"context"
	"log"
	pb "statsServer/synchProto"
	"testing"

	grpc "google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

// TestSendStats sends the stats for one light bulb
func TestSendStats(t *testing.T) {
	log.Println("starting TestSendStats")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSynchronizerClient(conn)
	d := pb.Device{
		Building: 1,
		Room:     13,
		Label:    "Front",
		LedOn:    false,
		OnTime:   21,
	}

	stream, err := c.Synch(context.Background(), &d)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		feature, _ := stream.Recv()
		if feature != nil {
			log.Println(feature)
		}
	}
}

func TestSwitchLight(t *testing.T) {
	log.Println("starting TestSwitchLight")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSynchronizerClient(conn)
	d := pb.Device{
		Building: 1,
		Room:     13,
		Label:    "Front",
		LedOn:    true,
		OnTime:   21,
	}
	c.LightSwitcher(context.Background(), &d)

}
