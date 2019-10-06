package entities

import (
	"internal/app/model"

	"github.com/jinzhu/gorm"
)

type ClientEntity struct {
	gorm.Model
	model.Client
}

func NewClientEntity(model *model.Client) *ClientEntity {
	clientEntity := new(ClientEntity)
	clientEntity.Client = *model
	return clientEntity
}

func NewClient(entity ClientEntity) *model.Client {
	model := model.NewUser(entity.Email, entity.Name)
	return model
}
