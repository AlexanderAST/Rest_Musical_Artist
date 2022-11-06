package main

import (
	"Music_Rest_Api/pkg/handler"
	"Music_Rest_Api/pkg/store"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	log.Print("Server start")
	route := gin.Default()
	store.ConnectDB()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello Music World!"})

	})
	route.GET("/tracks", handler.GetAllTrack)
	route.POST("/tracks", handler.CreateTrack)
	route.GET("/tracks/:id", handler.GetTrack)
	route.PATCH("/tracks/:id", handler.UpdateTrack)
	route.DELETE("/tracks/:id", handler.Delete)

	route.Run()
}
