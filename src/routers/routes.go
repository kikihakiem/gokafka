package routers

import (
	"controllers"
	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine) {
	engine.GET("/prime/*number", controllers.Prime)
	engine.GET("/fibonacci/*number", controllers.Fibonacci)
}
