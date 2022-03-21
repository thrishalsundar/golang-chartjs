package services

import "dhlabs/backend/models"

type TransactService interface {
	NewTrans(*models.Transact) error
	GetTrans() ([]*models.Transact, error)
}
