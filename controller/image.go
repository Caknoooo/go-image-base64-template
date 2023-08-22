package controller

import (
	"golang-base64-file-encryption-template/dto"
	"golang-base64-file-encryption-template/service"
	"os"

	"github.com/gin-gonic/gin"
)

type (
	ImageController interface{
		UploadImage(ctx *gin.Context)
		GetImage(ctx *gin.Context)
	}

	imageController struct {
		imageService service.ImageService
	}
)

const PATH = "storage/"

func NewImageController(is service.ImageService) ImageController {
	return &imageController{
		imageService: is,
	}
}

func (ic *imageController) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	req := dto.ImageRequest{
		Image: file,
	}

	res, err := ic.imageService.UploadImage(ctx, req)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    res,
	})
}

func (ic *imageController) GetImage(ctx *gin.Context) {
	id := ctx.Param("id")

	imagePath := PATH + id

	_, err := os.Stat(imagePath)
	if os.IsNotExist(err) {
		ctx.JSON(400, gin.H{
			"message": "image not found",
		})
		return
	}

	ctx.File(imagePath)
}