package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"todopoint/common2/d7redis"
	"todopoint/member/api"
	"todopoint/member/api/controller"
	"todopoint/member/out/ent"
	"todopoint/member/out/persistence"
	"todopoint/member/service"
	"todopoint/member/utils/config"
)

func main() {
	// Init DB
	client, err := config.NewEntClient(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	if err != nil {
		log.Println("Fail to initialize client")
	}

	// set client
	config.SetClient(client)

	// Redis
	redisClient := d7redis.NewRedisClient()
	d7redis.SetClient(redisClient)

	ok := d7redis.IsExist(redisClient)
	if !ok {
		panic("Redis client did not set")
	}

	// Set Router
	store := persistence.NewStore()
	memService := service.NewMemberService(store)
	memController := controller.NewMemberController(*memService)
	routes := api.NewMemberRouter(memController)
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
