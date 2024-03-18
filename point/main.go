package main

import (
	"log"
	"os"
	"path/filepath"
	"todopoint/common/db/config"
)

func main() {
	// Init DB
	client, err := config.NewEntClient(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	if err != nil {
		log.Println("Fail to initialize client")
	}
	// set client
	config.SetClient(client)

	//
}
