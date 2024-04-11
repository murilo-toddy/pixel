package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base      `valid:"required"`
    CreatedAt time.Time  `json:"created_at" gorm:"type:varchar(20)" valid:"notnull"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"type:varchar(255) "valid:"notnull"`
    Accounts  []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	return err
}

func NewBank(code, name string) (*Bank, error) {
	bank := Bank{
        Code: code,
        Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
