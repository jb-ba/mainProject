package main

import (
	"log"

	coap "github.com/go-ocf/go-coap"
)

func handleA(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got message in handleA: path=%q: %#v from %v", req.Msg.Path(), req.Msg, req.Client.RemoteAddr())
	w.SetContentFormat(coap.TextPlain)
	log.Printf("Transmitting from A")
	if _, err := w.Write([]byte("hello world")); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}

func handleB(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got message in handleB: path=%q: %#v from %v", req.Msg.Path(), req.Msg.PathString(), req.Client.RemoteAddr())
	resp := w.NewResponse(coap.Content)
	resp.SetOption(coap.ContentFormat, coap.TextPlain)
	resp.SetPayload([]byte("good bye!"))
	log.Printf("Transmitting from B %#v", resp)
	if err := w.WriteMsg(resp); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}

func handleC(w coap.ResponseWriter, req *coap.Request) {
	log.Printf("Got message in handleC: path=%q: %#v from %v", req.Msg.Path(), string(req.Msg.Payload()), req.Client.RemoteAddr())
	resp := w.NewResponse(coap.Content)
	resp.SetOption(coap.ContentFormat, coap.TextPlain)
	resp.SetPayload([]byte("good bye!"))
	log.Printf("Transmitting from C %#v", resp)
	if err := w.WriteMsg(resp); err != nil {
		log.Printf("Cannot send response: %v", err)
	}
}

func main() {
	mux := coap.NewServeMux()
	mux.Handle("/a", coap.HandlerFunc(handleA))
	mux.Handle("/b", coap.HandlerFunc(handleB))
	mux.Handle("/c", coap.HandlerFunc(handleC))

	log.Fatal(coap.ListenAndServe("udp", ":5688", mux))
}
