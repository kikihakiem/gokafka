package main

import (
	"github.com/gin-gonic/gin"
	"routers"
)

func main() {
	engine := gin.Default()
	routers.InitRoutes(engine)
	engine.Run() // listen and serve on 0.0.0.0:8080
}
