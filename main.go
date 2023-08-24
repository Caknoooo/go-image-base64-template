package main

import (
	"golang-base64-file-encryption-template/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.POST("/upload", controller.UploadImage)
	server.GET("/image/storage/:id", controller.GetImage)

	server.Run(":8080")
}