package main

import (
	"github.com/RyomaK/cancellation/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/", routes.Index)
	router.GET("/login", routes.Login)
	router.POST("/login",routes.Login)

	router.Run(":8080")

}
