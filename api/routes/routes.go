package routes

import (
	"github.com/diegolopezcode/go-hmtx-server/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	app *gin.Engine,
) {
	v1 := app.Group("/api/v1")
	{
		PublicRoutes(v1)
	}

}

func PublicRoutes(r *gin.RouterGroup) {
	r.GET("/", handler.GetCounter)
	r.POST("/increment", handler.IncrementCounter)
	r.POST("/decrement", handler.DecrementCounter)
}
