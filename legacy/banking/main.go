package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todopoint/common/db/config"

	"todopoint/banking/controller"
	"todopoint/banking/repo"
	"todopoint/banking/router"
	"todopoint/banking/service"
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
	store := repo.NewStore()
	accountService := service.NewBankAccountService(store)
	accountController := controller.NewBankAccountController(*accountService)
	routes := router.NewRouter(accountController)
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
