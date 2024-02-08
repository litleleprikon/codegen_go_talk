package main

import (
	"log"
	"net"

	"github.com/labstack/echo/v4"
	grpcGen "github.com/litleleprikon/codegen_go_talk/pkg/api/grpc/coverage/v1"
	v1 "github.com/litleleprikon/codegen_go_talk/pkg/api/rest/coverage/v1"
	coverageGRPCServer "github.com/litleleprikon/codegen_go_talk/pkg/server/grpc"
	coverageRESTServer "github.com/litleleprikon/codegen_go_talk/pkg/server/rest"
	"github.com/litleleprikon/codegen_go_talk/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	repo := repository.New()
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	covGRPCServer := coverageGRPCServer.New(repo)
	grpcGen.RegisterCoverageServiceServer(grpcServer, covGRPCServer)
	reflection.Register(grpcServer)

	var covRESTServer = coverageRESTServer.New(repo)
	e := echo.New()
	v1.RegisterHandlersWithBaseURL(e, covRESTServer, "/api/v1")

	log.Println("Starting...")
	go grpcServer.Serve(lis)
	e.Logger.Fatal(e.Start("localhost:8080"))
}
