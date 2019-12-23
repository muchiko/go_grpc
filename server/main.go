package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/muchiko/go_grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	pb.RegisterSocketServiceServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

type server struct{}

func (s *server) Transport(stream pb.SocketService_TransportServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println(in.Message)

		err = stream.Send(&pb.Payload{
			Message: in.Message,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
