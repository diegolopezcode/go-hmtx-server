package main

import (
	"github.com/diegolopezcode/go-hmtx-server/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	routes.SetupRoutes(app)
	app.Run(":8080")
}
