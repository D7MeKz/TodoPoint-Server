package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todopoint/common/db/config"
	"todopoint/member/out/persistence"
	"todopoint/member/router"
	"todopoint/member/router/controller"
	"todopoint/member/service"
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
	// Set Router
	store := persistence.NewStore()
	memService := service.NewMemberService(store)
	memController := controller.NewMemberController(*memService)
	routes := router.NewMemberRouter(memController)
	server := &http.Server{
		Addr:           ":3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server started on port 3000")
	log.Fatal(server.ListenAndServe())

}
