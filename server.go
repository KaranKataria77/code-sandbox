package main

import (
	pb "code-execution-sandbox/proto"
	"fmt"
	"log"
	"net"
	"net/http"

	"code-execution-sandbox/internal/setup"

	"github.com/soheilhy/cmux"

	"code-execution-sandbox/internal/routes"

	"google.golang.org/grpc"
)

// type server struct {
// 	pb.UnimplementedCodeExecutionServiceServer
// }

// func (s *server) ExecuteCode(ctx context.Context, req *pb.ExecutionRequest) (*pb.ExecutionResponse, error) {
// 	lang := req.GetLanguage()
// 	code := req.GetCode()

// 	var cmd *exec.Cmd

// 	switch lang {
// 	case "python":
// 		cmd = exec.Command("python3", "-c", code)
// 	case "javascript":
// 		cmd = exec.Command("node", "-e", code)
// 	default:
// 		return &pb.ExecutionResponse{Error: "Unsupported language"}, nil
// 	}

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		return &pb.ExecutionResponse{Error: "Error while running code " + err.Error()}, nil
// 	}

// 	return &pb.ExecutionResponse{Output: string(output)}, nil
// }

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen " + err.Error())
	}

	// create a connection multiplexer
	m := cmux.New(listener)

	grpcl := m.Match(cmux.HTTP2())
	httpl := m.Match(cmux.HTTP1Fast())

	grpcServer := grpc.NewServer()
	pb.RegisterFileDownloadServiceServer(grpcServer, &setup.Server{})
	// pb.RegisterCodeExecutionServiceServer(grpcServer, &server{})

	httpServer := &http.Server{
		Handler: routes.SetupRoutes(),
	}

	go grpcServer.Serve(grpcl)
	go httpServer.Serve(httpl)

	fmt.Println("multiplexer server is running")

	if err := m.Serve(); err != nil {
		log.Println("Error while running CMUX " + err.Error())
	}
}
