package main

import (
	"log"
	"modules/d7mysql"
	"modules/d7mysql/ent"
	"net/http"
	"time"
	"todopoint/user/api"
	"todopoint/user/api/controller"
	"todopoint/user/service"
	"todopoint/user/spi/mysqlS"
)

func main() {
	// Database setup
	// - MySql(Ent)
	client, err := d7mysql.NewEntClient(d7mysql.WithEnv("./.env"))

	if err != nil {
		log.Printf("err : %s", err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	// Set Router
	store := mysqlS.NewProfileStore(client)
	svc := service.NewUserService(store)
	ctr := controller.NewUserController(svc)
	routes := api.NewUserRouter(ctr)
	srv := http.Server{
		Addr:           ":3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Server started on port 3000")
	log.Fatal(srv.ListenAndServe())
}
