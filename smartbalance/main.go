package main

import (
	"fmt"
	"net"
	"smb/app/handler"
	"smb/app/repository"
	"smb/pkg/api"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)


type GRPCserver struct {
	api.UnimplementedSmartBalanceServiceServer
}

func init(){
	repository.NewPangolinDB(repository.Config{
		Host: "172.26.0.2",
		Port: "5433",
		Username: "postgres",
		Password: "secret",
		DBName: "smartbalance",
		SSLMode: "disable",
	})
}


	
func main(){

	log.SetFormatter(new(log.JSONFormatter))

	//Start gRPC server 
	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		fmt.Println(err)
	}

	s := grpc.NewServer()
	grpc_srv := GRPCserver{}

	api.RegisterSmartBalanceServiceServer(s, grpc_srv)

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve : %v", err)
	}
		

	// Start Gin Server
	hanlers := new(handler.Handler)
	fmt.Println("Starting Gin server...")
	srv := new(handler.Server)
	if err := srv.Run("8080", hanlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server:% s", err.Error())
	}
}

