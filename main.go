package main

import (
	controllers "api/controllers"

	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()

	go func(){router.POST("/user", controllers.PostUser)}()
	go func(){router.POST("/auth", controllers.Authentication)}()
	go func(){router.POST("/",controllers.Home)}()

	router.Run("localhost:9090")
}

