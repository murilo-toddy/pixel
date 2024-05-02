package model

import (
	"encoding/json"

	"github.com/asaskevich/govalidator"
)

type Transaction struct {
    ID string `json:"id" validate:"required,uuid4"`
    AccountID string `json:"account_id" validate:"required,uuid4"`
    Amount float64 `json:"account" validate:"required,numeric"`
    PixKeyTo string `json:"pix_key_to" validate:"required"`
    PixKeyKindTo string `json:"pix_key_kind_to" validate:"required"`
    Description string `json:"description" validate:"required"`
    Status string `json:"status" validate:"required"`
    Error string `json:"error" validate:"required"`
}

func (t *Transaction) validate() error {
    _, err := govalidator.ValidateStruct(t)
    return err
}

func (t *Transaction) ParseJson(data []byte) error {
    err := json.Unmarshal(data, t)
    if err != nil {
        return err
    }
    return t.validate()
}
