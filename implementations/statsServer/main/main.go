package main

import (
	"context"
	"flag"
	"log"
	"net"
	pb "statsServer/synchProto"
	"strconv"
	"strings"

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
	in        *pb.Device
	streamOut *pb.Synchronizer_SynchServer
}

func (s *synchServer) Synch(in *pb.Device, streamOut pb.Synchronizer_SynchServer) error {
	id := createID(in.GetBuilding(), in.GetRoom(), in.Label)
	ss := streamSlave{
		id:        id,
		streamOut: &streamOut,
	}
	slaves[&ss] = make(chan bool)
	log.Println(in)
	for {
		select {
		case on := <-slaves[&ss]:
			log.Println(on)
			in.LedOn = on
			streamOut.Send(in)
		}
	}
	return nil
}
func (s *synchServer) LightSwitcher(ctx context.Context, in *pb.Device) (*pb.Empty, error) {
	slaves[checkSlaves(createID(in.Building, in.Room, in.Label))] <- in.LedOn
	return &pb.Empty{}, nil
}

func splitID(id string) (int, int, string) {
	s := strings.Split(id, "-")
	log.Println(s)
	b, _ := strconv.Atoi(s[0])
	r, _ := strconv.Atoi(s[1])
	return b, r, s[2]
}

func createID(building int32, room int32, label string) string {
	log.Println(strconv.Itoa(int(building)) + "-" + strconv.Itoa(int(room)) + "-" + label)
	return strconv.Itoa(int(building)) + "-" + strconv.Itoa(int(room)) + "-" + label
}

func checkSlaves(id string) *streamSlave {
	for slave, _ := range slaves {
		if slave.id == id {
			return slave
		}
	}
	return nil
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
