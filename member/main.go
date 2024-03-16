package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todopoint/common/db/config"
	"todopoint/member/internal/middleware"
	"todopoint/member/internal/router"
)

func main() {
	// init Ent Client
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

	// init mux
	r := mux.NewRouter()
	r.Use(middleware.Header) // Middleware header
	router.RegisterMainRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server started on port 3000")
	log.Fatal(srv.ListenAndServe())
}
