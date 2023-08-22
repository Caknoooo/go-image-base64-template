package dto

import (
	"mime/multipart"
)

type (
	ImageRequest struct {
		Image *multipart.FileHeader `json:"image" form:"image"`
	}

	ImageResponse struct {
		ID       string  `json:"id"`
		Filename string `json:"filename"`
		Path     string `json:"path"`
	}
)
