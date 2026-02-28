package main

import (
	"fmt"

	"go-event-mgmt-app/database"
	"go-event-mgmt-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	fmt.Println("Server Started")
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8081")
}
