package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	coap "github.com/go-ocf/go-coap"
	proto "github.com/golang/protobuf/proto"
)

var (
	btnMsgs          []*ButtonMessage
	lastLedState     bool
	lastTurnOnTime   time.Time
	currentOnTimeSec int32
)

func observeLed(w coap.ResponseWriter, req *coap.Request) {

	log.Printf("Got message path=%q: %#v from %v", req.Msg.Path(), req.Msg, req.Client.RemoteAddr())
	go periodicTransmitter(w, req)
}

func periodicTransmitter(w coap.ResponseWriter, req *coap.Request) {
	subded := time.Now()
	for {
		nextTime := time.Now().Truncate(time.Second)
		nextTime = nextTime.Add(time.Second)
		time.Sleep(time.Until(nextTime))
		err := sendResponse(w, req, subded)
		if err != nil {
			log.Printf("Error on transmitter, stopping: %v", err)
			return
		}
		time.Sleep(time.Second)
	}
}

func sendResponse(w coap.ResponseWriter, req *coap.Request, subded time.Time) error {
	log.Println("sendRespsonse observable")
	resp := w.NewResponse(coap.Content)
	resp.SetOption(coap.ContentFormat, coap.TextPlain)
	resp.SetPayload([]byte(fmt.Sprintf("Been running for %v", time.Since(subded))))
	return w.WriteMsg(resp)
}

func handleClick(w coap.ResponseWriter, req *coap.Request) {
	// log.Printf("Got message in click: path=%q: %#v from %v", req.Msg.Path(), string(req.Msg.Payload()), req.Client.RemoteAddr())
	// log.Printf("The payload is %#v", req.Msg.Payload())
	msg := &ButtonMessage{}
	err := proto.Unmarshal(req.Msg.Payload(), msg)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	lastLedState = *msg.LedOn
	if lastLedState {
		lastTurnOnTime = time.Now()
	} else {
		currentOnTimeSec += int32(time.Now().Sub(lastTurnOnTime).Seconds())
	}
	// respond(&w)
	log.Println(strconv.Itoa(int(*msg.Building)) + " ; " + strconv.Itoa(int(*msg.Room)) + " ; " + *msg.Label)
}

func respond(wi *coap.ResponseWriter) {
	w := *wi
	resp := w.NewResponse(coap.Content)
	resp.SetOption(coap.ContentFormat, coap.TextPlain)
	resp.SetPayload([]byte("Received your stuff!"))
	log.Println()
	log.Printf("Transmitting from D %#v", resp)
	if err := w.WriteMsg(resp); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}

// uploads data every hour to the server.
func uploadTicker() {
	for {
		nextTime := time.Now().Truncate(time.Minute)
		nextTime = nextTime.Add(time.Minute)
		time.Sleep(time.Until(nextTime))
		if lastLedState {
			currentOnTimeSec += int32(time.Now().Sub(lastTurnOnTime).Seconds())
			lastTurnOnTime = time.Now()
		}
		sendToServer()
		currentOnTimeSec = 0
	}
}

func sendToServer() {
	log.Println(currentOnTimeSec)
}

func main() {
	mux := coap.NewServeMux()
	mux.Handle("/clicked", coap.HandlerFunc(handleClick))
	mux.Handle("/observe", coap.HandlerFunc(observeLed))
	go uploadTicker()
	log.Fatal(coap.ListenAndServe("udp", ":5688", mux))
}
