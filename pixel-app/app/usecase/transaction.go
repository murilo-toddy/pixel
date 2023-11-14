package usecase

import (
	"github.com/murilo-toddy/pixel/domain/model"
	"github.com/prometheus/common/log"
)

type TransactionUseCase struct {
    TransactionReponsitory model.TransactionRepositoryInterface
    PixKeyRepository model.PixKeyRepositoryInterface
}

func (t *TransactionUseCase) Register(
    accountID string, 
    amount float64, 
    pixKeyTo string, 
    pixKeyKindTo string, 
    description string,
) (*model.Transaction, error) {
    account, err := t.PixKeyRepository.FindAccount(accountID)
    if err != nil {
        return nil, err
    }

    pixKey, err := t.PixKeyRepository.FindKeyByID(pixKeyTo, pixKeyKindTo)
    if err != nil {
        return nil, err
    }

    transaction, err := model.NewTransaction(account, amount, pixKey, description)
    if err != nil {
        return nil, err
    }

    err = t.TransactionReponsitory.Save(transaction)
    if err != nil {
        return nil, err
    }

    return transaction, nil
}

func (t *TransactionUseCase) Confirm(transactionID string) (*model.Transaction, error) {
    transaction, err := t.TransactionReponsitory.Find(transactionID)
    if err != nil {
        return nil, err
    }

    err = transaction.Confirm()
    if err != nil {
        return nil, err
    }
    
    err = t.TransactionReponsitory.Save(transaction)
    if err != nil {
        return nil, err
    }
    
    return transaction, nil
}


func (t *TransactionUseCase) Complete(transactionID string) (*model.Transaction, error) {
    transaction, err := t.TransactionReponsitory.Find(transactionID)
    if err != nil {
        return nil, err
    }

    err = transaction.Complete()
    if err != nil {
        return nil, err
    }

    err = t.TransactionReponsitory.Save(transaction)
    if err != nil {
        return nil, err
    }
    
    return transaction, nil
}

func (t *TransactionUseCase) Error(transactionID string, reason string) (*model.Transaction, error) {
    transaction, err := t.TransactionReponsitory.Find(transactionID)
    if err != nil {
        return nil, err
    }

    err = transaction.Cancel(reason)
    if err != nil {
        return nil, err
    }

    err = t.TransactionReponsitory.Save(transaction)
    if err != nil {
        return nil, err
    }
    
    return transaction, nil
}
