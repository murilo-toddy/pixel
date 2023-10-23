package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
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
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
}

func (t *Transaction) isValid() error {
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
		AccountFrom: accountFrom,
		Amount:      amount,
		PixKeyTo:    pixKeyTo,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (t *Transaction) Complete() error {
	if err := t.isValid(); err != nil {
		return err
	}
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Confirm() error {
	if err := t.isValid(); err != nil {
		return err
	}
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Transaction) Cancel(description string) error {
	if err := t.isValid(); err != nil {
		return err
	}
	t.Status = TransactionError
	t.CancelDescription = description
	t.UpdatedAt = time.Now()
	return nil
}
