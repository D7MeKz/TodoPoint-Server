package main

import (
	"log"
	"modules/v2/d7mongo"
	"net/http"
	"time"
	"todopoint/image/api"
	"todopoint/image/service"
	"todopoint/image/spi"
)

func main() {
	// mongo client
	client := d7mongo.NewMongoClient(d7mongo.WithEnv(".env"))

	// Set router
	store := spi.NewImageStore(client.Client)
	svc := service.NewImageService(store)
	ctr := api.NewImageController(svc)
	routes := api.NewImageRouter(ctr)
	server := &http.Server{
		Addr:           ":3003",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server started on port 3003")
	log.Fatal(server.ListenAndServe())
}
