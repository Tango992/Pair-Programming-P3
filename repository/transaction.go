package repository

import "pair-programming/models"

type Transaction interface {
	Post(*models.Transaction) error
	GetAll() ([]models.Transaction, error)
	GetById(string) (models.Transaction, error)
	PutById(string, *models.Transaction) error
	DeleteById(string) error
}