package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"modules/v2/d7mysql/ent"
	"net/http"
	"os"
	"time"
	"todopoint/user/v2/api"
	"todopoint/user/v2/service"
	"todopoint/user/v2/spi"
)

func main() {
	// Set logger
	// logger := logging.NewAppLogger()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	client, err := ent.Open("mysql", dsn, ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Println(dsn)
			fmt.Print("\n")
		}
	}))
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	//// Database setup
	//client, err := d7mysql.NewEntClient()
	//if err != nil {
	//	logger.Fatal(nil, 0, err.Error())
	//}
	//defer func(client *ent.Client) {
	//	_ = client.Close()
	//}(client)
	//

	mysqlStore := spi.NewUserStore(client)
	svc := service.NewUserService(mysqlStore)
	ctr := api.NewUserController(svc)
	routes := api.NewUserRouter(ctr)

	srv := http.Server{
		Addr:           ":3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(srv.ListenAndServe())

}
