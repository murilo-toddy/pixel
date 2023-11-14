package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/murilo-toddy/pixel/domain/model"
)

type PixKeyRepositoryDB struct {
	DB *gorm.DB
}

func (r *PixKeyRepositoryDB) AddBank(bank *model.Bank) error {
	err := r.DB.Create(bank).Error
	return err
}

func (r *PixKeyRepositoryDB) AddAccount(account *model.Account) error {
	err := r.DB.Create(account).Error
	return err
}

func (r *PixKeyRepositoryDB) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.DB.Create(pixKey).Error
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (r *PixKeyRepositoryDB) FindKeyByID(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey
	r.DB.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("No pixkey was found for key=<%s>, kind=<%s>", key, kind)
	}
	return &pixKey, nil
}

func (r *PixKeyRepositoryDB) FindAccount(id string) (*model.Account, error) {
	var account model.Account
	r.DB.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("No account was found for id=<%s>", id)
	}
	return &account, nil
}

func (r *PixKeyRepositoryDB) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank
	r.DB.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("No bank was found for id=<%s>", id)
	}
	return &bank, nil
}
