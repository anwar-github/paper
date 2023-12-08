package service

import (
	"errors"
	"log"
	"math/rand"
	"papper/exception"
	"papper/model"
	"reflect"
	"time"
)

type DisbursementRequest struct {
	UserCode      string `json:"user_code" validate:"required"`
	RequestId     string `json:"request_id" validate:"required"`
	AccountNumber string `json:"account_number" validate:"required"`
	Provider      string `json:"provider" validate:"required"`
	Amount        int64  `json:"amount" validate:"required"`
}

type DisbursementResponse struct {
	Request      *DisbursementRequest `json:"request"`
	Wallet       *model.Wallet        `json:"wallet"`
	Disbursement *model.Disbursement  `json:"disbursement"`
}

func (service Service) Disbursement(request *DisbursementRequest) (*DisbursementResponse, error) {
	var (
		response DisbursementResponse
	)

	response.Request = request

	// init database transaction
	// =================================================================================================================
	tx, err := service.Repository.Transaction()
	if err != nil {
		return handleError(response, err)
	}
	defer service.Repository.Rollback(tx)

	// validate disbursement
	// =================================================================================================================
	disbursement, err := service.Repository.FindDisbursementByRequestID(tx, request.RequestId)
	if err != nil {
		return handleError(response, err)
	}
	response.Disbursement = disbursement
	if err := disbursementValidation(disbursement, request.RequestId); err != nil {
		return handleError(response, err)
	}

	// validate wallet
	// =================================================================================================================
	wallet, err := service.Repository.FindWalletForUpdateByUserCode(tx, request.UserCode)
	if err != nil {
		return handleError(response, err)
	}
	response.Wallet = wallet
	if err := service.WalletValidation(wallet, request.Amount); err != nil {
		return handleError(response, err)
	}
	// deduct wallet
	// =================================================================================================================
	newBalance := wallet.Balance - request.Amount
	if err := service.Repository.UpdateWalletBalanceByUserCode(
		tx,
		newBalance,
		request.UserCode,
	); err != nil {
		return handleError(response, err)
	}
	wallet.Balance = newBalance

	// store disbursement
	// =================================================================================================================
	orderID := generateOrderID(15)
	disbursement, err = service.Repository.InsertDisbursement(tx, &model.Disbursement{
		Order:         orderID,
		Success:       true,
		UserCode:      request.UserCode,
		AccountNumber: request.AccountNumber,
		Amount:        request.Amount,
		Balance:       newBalance,
		Provider:      request.Provider,
		Request:       request.RequestId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	})
	if err != nil {
		return handleError(response, err)
	}
	response.Disbursement = disbursement

	// end database transaction
	// =================================================================================================================
	if err = service.Repository.Commit(tx); err != nil {
		return handleError(response, err)
	}

	return handleSuccess(response)
}

func handleError(response DisbursementResponse, err error) (*DisbursementResponse, error) {
	log.Print(err)
	return &response, err
}

func handleSuccess(response DisbursementResponse) (*DisbursementResponse, error) {
	return &response, nil
}

func (service Service) WalletValidation(wallet *model.Wallet, amount int64) error {
	if reflect.ValueOf(wallet).IsNil() {
		return errors.New(exception.InsufficientBalance)
	}
	if wallet.Balance < amount {
		return errors.New(exception.InsufficientBalance)
	}
	return nil
}

func disbursementValidation(disbursement *model.Disbursement, requestID string) error {
	if !reflect.ValueOf(disbursement).IsNil() {
		return errors.New(exception.DuplicateRequest)
	}

	return nil
}

func generateOrderID(length int) string {
	// ... create character set ...
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// ... generate random string ...
	result := make([]byte, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}
	return string(result)
}
