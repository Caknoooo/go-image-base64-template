package entities

import "github.com/google/uuid"

type (
	Image struct {
		ID       uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
		Filename string    `json:"filename"`
		Path     string    `json:"path"`
	}
)
