package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	coap "github.com/go-ocf/go-coap"
)

func TestClient(*testing.T) {
	log.Println("testing")
	// co, err := coap.Dial("udp", "195.249.187.164:30002")
	co, err := coap.Dial("udp", "10.3.141.225:30002")
	// co, err := coap.Dial("udp", "0.0.0.0:5688")
	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	path := "/c"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	resp, err := co.Post(path, coap.TextPlain, bytes.NewReader([]byte("timesHallo")))

	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	log.Printf("Response payload: %v", string(resp.Payload()))
}
