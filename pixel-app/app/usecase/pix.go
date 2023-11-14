package usecase

import "github.com/murilo-toddy/pixel/domain/model"

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountID string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountID)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, kind, key)
	if err != nil {
		return nil, err
	}

	_, err = p.PixKeyRepository.RegisterKey(pixKey)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	return p.PixKeyRepository.FindKeyByID(key, kind)
}
