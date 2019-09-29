package entities

import (
	"internal/app/model"

	"github.com/jinzhu/gorm"
)

type ClientEntity struct {
	model.Client
	gorm.Model
}
