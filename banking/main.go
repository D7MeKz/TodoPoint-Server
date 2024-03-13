package main

import (
	"github.com/gin-gonic/gin"
	"todopoint/banking/router"
)

func main() {
	// Init DB

	// Set Router
	r := gin.Default()
	router.NewRouter(r)
	r.Run(":3001")
}
