package server

import (
	"fmt"
	"net"
	"server-streaming-rpc/pb"
	"time"

	"google.golang.org/grpc"
)

type StatusServer struct {
	pb.UnimplementedStatusServiceServer
}

func (s *StatusServer) StreamStatus(req *pb.StreamRequest, stream grpc.ServerStreamingServer[pb.StreamResponse]) error {

	taskId := req.TaskId
	fmt.Printf("Received request for task: %s\n", taskId)

	//simula um download de um arquivo
	for i := 0; i <= 100; i += 5 {
		status := &pb.StreamResponse{
			Message:  "Progressing",
			Progress: int64(i),
		}
		err := stream.Send(status)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)

	}
	return nil
}

func Run() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStatusServiceServer(grpcServer, &StatusServer{})
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}

}
