package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByID(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
	FindBank(id string) (*Bank, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" gorm:"type:varchar(20)" valid:"notnull"`
	Key       string   `json:"key" gorm:"type:varchar(255)" valid:"notnull"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
}

func (pixKey *PixKey) validate() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("Invalid key type")
	}

	if pixKey.Kind != "active" && pixKey.Kind != "inactive" {
		return errors.New("Invalid key status")
	}

	return err
}

func NewPixKey(account *Account, kind, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:      kind,
		Key:       key,
		Account:   account,
		AccountID: account.ID,
		Status:    "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	pixKey.UpdatedAt = time.Now()

	err := pixKey.validate()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
