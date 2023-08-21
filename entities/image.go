package entities

type (
	Image struct {
		ID       int64  `gorm:"primary_key;auto_increment" json:"id"`
		Filename string `json:"filename"`
		Path     string `json:"path"`
	}
)
