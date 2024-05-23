package main

import (
	"fmt"
	"smb/app/handler"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)
	
func main(){

	// Start Gin Server
	hanlers := new(handler.Handler)
	fmt.Println("Starting Gin server...")
	srv := new(handler.Server)
	if err := srv.Run("8080", hanlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server:% s", err.Error())
	}
}

