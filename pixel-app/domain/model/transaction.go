package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionError     string = "error"
	TransactionConfirmed string = "confirmed"
)

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid;" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixKeyToID        string   `gorm:"column:pix_key_to_id;type:uuid;" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"-"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (t *Transaction) validate() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("Amount must be a positive number")
	}
	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError && t.Status != TransactionConfirmed {
		return errors.New("Invalid transaction status")
	}
	return err
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AccountFrom:   accountFrom,
		AccountFromID: accountFrom.ID,
		Amount:        amount,
		PixKeyTo:      pixKeyTo,
		PixKeyToID:    pixKeyTo.ID,
		Status:        TransactionPending,
		Description:   description,
	}
	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()
	err := transaction.validate()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (t *Transaction) Complete() error {
	if err := t.validate(); err != nil {
		return err
	}
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Confirm() error {
	if err := t.validate(); err != nil {
		return err
	}
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Cancel(description string) error {
	if err := t.validate(); err != nil {
		return err
	}
	t.Status = TransactionError
	t.CancelDescription = description
	t.UpdatedAt = time.Now()
	return nil
}
