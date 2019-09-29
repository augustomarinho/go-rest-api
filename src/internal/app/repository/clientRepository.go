package repository

import "internal/app/model"

type ClientRepository interface {
	FindAll() ([]*model.Client, error)
	FindByEmail(email string) (*model.Client, error)
	Save(*model.Client) error
}

type ClientRepositoryImpl struct {
}

func (repository ClientRepositoryImpl) FindAll() ([]*model.Client, error) {
	return nil, nil
}

func (repository ClientRepositoryImpl) FindByEmail(email string) (*model.Client, error) {
	return nil, nil
}

func (repository ClientRepositoryImpl) Save(*model.Client) error {
	return nil, nil
}
