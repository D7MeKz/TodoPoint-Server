package main

import (
	"log"
	"modules/d7mysql/ent"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todopoint/auth/api"
	"todopoint/auth/api/controller"
	"todopoint/auth/config"
	"todopoint/auth/service"
	"todopoint/auth/spi/mysqlS"
	"todopoint/auth/spi/redisS"
)

func main() {
	// Database setup
	// - MySql(Ent)
	client, err := config.NewEntClient(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	// Set Router
	mysqlStore := mysqlS.NewUserStore(client)
	redisStore := redisS.NewRedisStore()
	svc := service.NewAuthService(redisStore, mysqlStore)
	ctr := controller.NewAuthController(svc)
	routes := api.NewAuthRouter(ctr)

	srv := http.Server{
		Addr:           ":3001",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server started on port 3000")
	log.Fatal(srv.ListenAndServe())

}
