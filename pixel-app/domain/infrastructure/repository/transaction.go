package repository

import (
	"fmt"

	"github.com/murilo-toddy/pixel/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDB struct {
	DB *gorm.DB
}

func (t *TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := t.DB.Create(transaction).Error
	return err
}

func (t *TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	err := t.DB.Save(transaction).Error
	return err
}

func (t *TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	t.DB.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was found for id=<%s>", id)
	}
	return &transaction, nil
}
