package test

import (
	"bytes"
	"log"
	"os"
	"testing"

	coap "github.com/go-ocf/go-coap"
)

func TestClient(*testing.T) {
	co, err := coap.Dial("udp", "10.3.141.1:5688")
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