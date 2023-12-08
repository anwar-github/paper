package service

import (
	"papper/model"
	"papper/repository"
)

type Service struct {
	Repository repository.Interface
}

type Interface interface {
	WalletValidation(wallet *model.Wallet, amount int64) error
	Disbursement(request *DisbursementRequest) (*DisbursementResponse, error)
}

func NewService(repository repository.Interface) Interface {
	return &Service{
		Repository: repository,
	}
}
