package repository

import (
	"errors"
	"internal/app/entities"
	"internal/app/infrastructure"
	"internal/app/model"
)

type ClientRepository interface {
	FindAll() ([]model.Client, error)
	FindByEmail(email string) (*model.Client, error)
	Save(*model.Client) error
}

type ClientRepositoryImpl struct {
	database *infrastructure.Database
}

func (repository *ClientRepositoryImpl) FindAll() ([]model.Client, error) {
	users := []entities.ClientEntity{}

	database := infrastructure.ConnectDatabase()
	err := database.GetConnection().Find(&users).Error

	if err != nil {
		return nil, errors.New("Problems to find all Clients on Database")
	}

	clients := make([]model.Client, len(users), cap(users))

	for index, element := range users {
		clients[index] = *entities.NewClient(element)
	}

	return clients, nil
}

func (repository *ClientRepositoryImpl) FindByEmail(email string) (*model.Client, error) {
	user := new(entities.ClientEntity)
	database := infrastructure.ConnectDatabase()
	userFound := database.GetConnection().Where("email=?", email).Find(user).RecordNotFound()

	if !userFound {
		client := entities.NewClient(*user)
		return client, nil
	}

	return nil, nil

}

func (repository *ClientRepositoryImpl) Save(model *model.Client) error {

	entity := entities.NewClientEntity(model)

	database := infrastructure.ConnectDatabase()
	err := database.GetConnection().Save(entity).Error

	if err != nil {
		return errors.New("Problems to find all Clients on Database")
	}

	return nil
}
