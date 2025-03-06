package routes

import (
	"code-execution-sandbox/internal/setup"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/test", setup.ExecuteCommand)
	return r
}
