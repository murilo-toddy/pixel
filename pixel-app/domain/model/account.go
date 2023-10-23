package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base    `valid:"required"`
	Owner   string    `json:"owner_name", valid:"notnull"`
	Bank    *Bank     `valid:"-"`
	Number  string    `json:"number" valid:"notnull"`
	PixKeys []*PixKey `valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	return err
}

func NewAccount(bank *Bank, number, owner string) (*Account, error) {
	account := Account{
		Owner:  owner,
		Bank:   bank,
		Number: number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}
	return &account, nil
}
