package test

import (
	"database/sql"
	"papper/model"
)

func (repository *MockRepository) FindWalletForUpdateByUserCode(tx *sql.Tx, userCode string) (*model.Wallet, error) {
	ret := repository.Called(tx, userCode)
	if _, ok := ret.Get(0).(*model.Wallet); ok {
		return ret.Get(0).(*model.Wallet), nil
	}
	return nil, nil
}

func (repository *MockRepository) UpdateWalletBalanceByUserCode(tx *sql.Tx, balance int64, userCode string) error {
	_ = repository.Called(tx, balance, userCode)
	return nil
}

func (repository *MockRepository) FindDisbursementByRequestID(tx *sql.Tx, requestID string) (*model.Disbursement, error) {
	ret := repository.Called(tx, requestID)
	if _, ok := ret.Get(0).(*model.Disbursement); ok {
		return ret.Get(0).(*model.Disbursement), nil
	}
	return nil, nil
}

func (repository *MockRepository) FindDisbursementByOrderID(tx *sql.Tx, orderID string) (*model.Disbursement, error) {
	ret := repository.Called(tx, orderID)
	if _, ok := ret.Get(0).(*model.Disbursement); ok {
		return ret.Get(0).(*model.Disbursement), nil
	}
	return nil, nil
}

func (repository *MockRepository) InsertDisbursement(tx *sql.Tx, disbursement *model.Disbursement) (*model.Disbursement, error) {
	ret := repository.Called(tx, disbursement)
	if _, ok := ret.Get(0).(*model.Disbursement); ok {
		return ret.Get(0).(*model.Disbursement), nil
	}
	return nil, nil
}

func (repository *MockRepository) Transaction() (*sql.Tx, error) {
	ret := repository.Called()
	if _, ok := ret.Get(0).(*sql.Tx); ok {
		return ret.Get(0).(*sql.Tx), nil
	}
	return nil, nil
}

func (repository *MockRepository) Rollback(tx *sql.Tx) error {
	_ = repository.Called(tx)
	return nil
}

func (repository *MockRepository) Commit(tx *sql.Tx) error {
	_ = repository.Called(tx)
	return nil
}
