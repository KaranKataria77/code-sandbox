package main

import (
	pb "code-sandbox/proto"
	"context"
	"fmt"
	"log"
	"net"
	"os/exec"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCodeExecutionServiceServer
}

func (s *server) ExecuteCode(ctx context.Context, req *pb.ExecutionRequest) (*pb.ExecutionResponse, error) {
	lang := req.GetLanguage()
	code := req.GetCode()

	var cmd *exec.Cmd

	switch lang {
	case "python":
		cmd = exec.Command("python3", "-c", code)
	case "javascript":
		cmd = exec.Command("node", "-c", code)
	default:
		return &pb.ExecutionResponse{Error: "Unsupported language"}, nil
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return &pb.ExecutionResponse{Error: "Error while running code " + err.Error()}, nil
	}

	return &pb.ExecutionResponse{Output: string(output)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen " + err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCodeExecutionServiceServer(grpcServer, &server{})

	fmt.Println("gRPC server is running on port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to server " + err.Error())
	}
}
