package main

import (
	"context"
	"flag"
	"log"
	"net"
	pb "statsServer/src/syncProto"
	"strconv"

	grpc "google.golang.org/grpc"
)

var (
	port   = flag.String("port", "3000", "Open port for connections.")
	slaves = make(map[*streamSlave](chan int))
)

type syncServer struct {
}

type streamSlave struct {
	id        string
	slave     *pb.Device
	streamOut *pb.Synchronizer_SyncServer
}

func (s *syncServer) Sync(in *pb.Device, streamOut pb.Synchronizer_SyncServer) error {
	ss := streamSlave{
		id:        createID(in.GetBuilding(), in.GetRoom(), in.Label),
		slave:     in,
		streamOut: &streamOut,
	}
	for s := range slaves {
		if s.id == ss.id {
			slaves[getSlave(s.id)] <- -1
		}
	}
	slaves[&ss] = make(chan int)

	for {
		select {
		case i := <-slaves[&ss]:
			if i == 0 || i == 1 {
				in.LedOn = i != 0
				streamOut.Send(in)
			} else if i == -1 {
				delete(slaves, &ss)
				return nil
			}
		default:
			if err := streamOut.Context().Err(); err != nil {
				delete(slaves, &ss)
				return nil
			}
		}
	}
}
func (s *syncServer) LightSwitcher(ctx context.Context, in *pb.Device) (*pb.Empty, error) {
	b := 0
	if in.LedOn {
		b = 1
	}
	if slaves[getSlave(createID(in.Building, in.Room, in.Label))] != nil {
		slaves[getSlave(createID(in.Building, in.Room, in.Label))] <- b
	}
	return &pb.Empty{}, nil
}

func (s *syncServer) ListDevices(in *pb.Empty, streamOut pb.Synchronizer_ListDevicesServer) error {
	log.Println(slaves)
	for sSlave := range slaves {
		streamOut.Send(sSlave.slave)
	}
	return nil
}

func createID(building int32, room int32, label string) string {
	return strconv.Itoa(int(building)) + "-" + strconv.Itoa(int(room)) + "-" + label
}

func getSlave(id string) *streamSlave {
	for slave := range slaves {
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
	pb.RegisterSynchronizerServer(grpcServer, &syncServer{})
	grpcServer.Serve(lis)
}
