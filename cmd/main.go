package main

import (
	"log"
	"github.com/hyperionXCF/rest-api/cmd/api"
)

// entry point for api

func main(){
	// create server instance
	server := api.NewApiServer(":8080",nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	} 
}