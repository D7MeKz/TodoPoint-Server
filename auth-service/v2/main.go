package main

import (
	"log"
	"modules/common/logging"
	"modules/d7mysql"
	"modules/d7mysql/ent"
	"modules/d7redis"
	"net/http"
	"time"
	"todopoint/auth/v2/api"
	"todopoint/auth/v2/service"
	"todopoint/auth/v2/spi/mysqlS"
	"todopoint/auth/v2/spi/redisS"
)

func main() {
	// Set logger
	logger := logging.NewAppLogger()

	// Database setup
	client, err := d7mysql.NewEntClient(d7mysql.WithEnv("./.env"))
	if err != nil {
		logger.Fatal(nil, 0, err.Error())
	}
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	rClient := d7redis.NewRedisClient()
	defer func(rClient *d7redis.RedisClient) {
		_ = rClient.Client.Close()
	}(rClient)

	// Set router and server
	mysqlStore := mysqlS.NewUserStore(client)
	redisStore := redisS.NewRedisStore(rClient.Client)
	svc := service.NewAuthService(redisStore, mysqlStore)
	ctr := api.NewAuthController(svc)
	routes := api.NewAuthRouter(ctr)

	srv := http.Server{
		Addr:           ":3001",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())

}
