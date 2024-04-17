package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todopoint/common/config/mongodb"
	"todopoint/task/in/router"
	"todopoint/task/in/router/controller"
	"todopoint/task/out/persistence"
	"todopoint/task/service"
)

func main() {
	// init MongoDB
	client := mongodb.NewMongoClient(filepath.Dir(os.Args[0]))
	mongodb.SetClient(client)

	store := persistence.NewTaskStore()
	taskService := service.NewTaskService(store)
	taskController := controller.NewTaskController(*taskService)
	routes := router.NewTaskRouter(taskController)
	server := &http.Server{
		Addr:           ":3001",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server started on port 3000")
	log.Fatal(server.ListenAndServe())
}
