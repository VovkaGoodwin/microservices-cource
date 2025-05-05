package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"user/internal/app"
	pb "user/proto"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	httpHandler := app.NewHttpHandler()
	go func() {
		err := httpHandler.ListenAndServe()
		if err != nil {
			log.Fatalf("http handler error: %v", err)
		}
	}()
	log.Println("start http server on :8080")

	grpcHandler := app.NewGrpcHandler()

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, grpcHandler)

	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	log.Println("grpc server started")
	sig := <-done
	log.Printf("stopping grpc server signal: %v", sig)

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.GracefulStop()
	log.Println("grpc server stopped")
}
