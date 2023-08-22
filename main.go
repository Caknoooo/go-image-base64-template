package main

import (
	"golang-base64-file-encryption-template/config"
	"golang-base64-file-encryption-template/controller"
	"golang-base64-file-encryption-template/repository"
	"golang-base64-file-encryption-template/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db *gorm.DB = config.SetupDatabaseConnection()
		imageRepository repository.ImageRepository = repository.NewImageRepository(db)
		imageService service.ImageService = service.NewImageService(imageRepository)
		imageController controller.ImageController = controller.NewImageController(imageService)
	)

	server := gin.Default()
	server.POST("/upload", imageController.UploadImage)
	server.GET("/image/storage/:id", imageController.GetImage)

	server.Run(":8080")
}