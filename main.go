package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/karansinghgit/golib/db"
	"github.com/karansinghgit/golib/routes"
)

func main() {
	// estabilish connection with the database
	db.Connect()

	router := gin.Default()
	routes.Routes(router)

	// listen to the router with port
	log.Fatal(router.Run("localhost:8080"))
}
