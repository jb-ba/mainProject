package main

import (
	"flag"
	"log"
	"net"
	pb "statsServer/synchProto"
	"strconv"
	"time"

	grpc "google.golang.org/grpc"
)

var (
	port   = flag.String("port", "3000", "Open port for connections.")
	slaves = make(map[*streamSlave](chan bool))
)

type synchServer struct {
}

type streamSlave struct {
	id        string
	streamOut *pb.Synchronizer_SynchServer
}

func (s *synchServer) Synch(in *pb.Device, streamOut pb.Synchronizer_SynchServer) error {
	id := strconv.Itoa(int(in.GetBuilding())) + "-" + strconv.Itoa(int(in.GetRoom())) + "-" + in.Label
	ss := streamSlave{
		id:        id,
		streamOut: &streamOut,
	}
	checkSlaves()
	slaves[&ss] = make(chan bool)
	go func() {
		time.Sleep(5 * time.Second)
		slaves[&ss] <- true
	}()
	go func() {
		time.Sleep(10 * time.Second)
		log.Println("wake up")
		slaves[&ss] <- false
	}()
	go func() {
		time.Sleep(30 * time.Second)
		log.Println("wake up")
		slaves[&ss] <- false
	}()
	for {
		select {
		case stat := <-slaves[&ss]:
			log.Println(stat)
			streamOut.Send(&pb.Device{
				Building: 1,
				Room:     2,
				Label:    "Back",
				LedOn:    stat,
				OnTime:   3,
			})
			// default:
			// 	log.Println("hallo")
		}
	}
	log.Println(&streamOut)
	log.Println(id)
	// time.Sleep(5 * time.Second)

	// if err := streamOut.Send(&pb.Device{
	// 	Building: 1,
	// 	Room:     2,
	// 	Label:    "back",
	// 	LedOn:    true,
	// 	OnTime:   3,
	// }); err != nil {
	// 	return err
	// }
	return nil
}

func checkSlaves() {
	for slave, _ := range slaves {
		if slave.id == "1-2-back" {
			log.Println(slave)
		}
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSynchronizerServer(grpcServer, &synchServer{})
	grpcServer.Serve(lis)
}
