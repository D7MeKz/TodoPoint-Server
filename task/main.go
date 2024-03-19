package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todopoint/common/db/config"
	"todopoint/task/in/router"
	"todopoint/task/in/web"
	"todopoint/task/out/persistence"
	"todopoint/task/service"
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
	taskStore := persistence.NewStore()
	taskService := service.NewTaskService(taskStore)
	taskController := web.NewTaskController(*taskService)
	taskRouterHandler := router.NewTaskRouter(taskController)
	server := &http.Server{
		Addr:           ":3001",
		Handler:        taskRouterHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server started on port 3001")
	log.Fatal(server.ListenAndServe())

}
