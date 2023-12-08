package test

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"papper/exception"
	"papper/model"
	"papper/service"
	"testing"
	"time"
)

var (
	mockRequest = &service.DisbursementRequest{
		UserCode:      "P1-GJSTCS",
		RequestId:     "merchant-000010",
		AccountNumber: "0627892217",
		Provider:      "BCA",
		Amount:        500,
	}
	mockWallet = &model.Wallet{
		UserCode:  "P1-GJSTCS",
		Balance:   1000,
		CreatedAt: "2023-12-07 20:00:07",
		UpdatedAt: "2023-12-07 20:00:07",
	}
	mockDisbursement = &model.Disbursement{
		Order:         "Order-000010",
		Success:       true,
		UserCode:      "P1-GJSTCS",
		AccountNumber: "0627892217",
		Amount:        500,
		Balance:       500,
		Provider:      "BCA",
		Request:       "merchant-000010",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	mockDisbursementResponse = &service.DisbursementResponse{
		Request:      mockRequest,
		Wallet:       mockWallet,
		Disbursement: mockDisbursement,
	}
	mockInsufficientRequest = &service.DisbursementRequest{
		UserCode:      "P1-GJSTCS",
		RequestId:     "merchant-000010",
		AccountNumber: "0627892217",
		Provider:      "BCA",
		Amount:        500000,
	}
)

func TestDisbursementService(t *testing.T) {
	t.Run("WALLET VALIDATION - NOT FOUND", func(t *testing.T) {
		// Call your function for test
		err := services.WalletValidation(nil, 1000)

		// Assert a result as expectations
		assert.Equal(t, err, errors.New(exception.InsufficientBalance))
	})
	t.Run("WALLET VALIDATION - INSUFFICIENT", func(t *testing.T) {
		// Call your function for test
		err := services.WalletValidation(mockWallet, 5000)

		// Assert a result as expectations
		assert.Equal(t, err, errors.New(exception.InsufficientBalance))
	})
	t.Run("WALLET VALIDATION - SUFFICIENT", func(t *testing.T) {
		// Call your function for test
		err := services.WalletValidation(mockWallet, 1)

		// Assert a result as expectations
		assert.Equal(t, err, nil)
	})
	t.Run("INSUFFICIENT BALANCE", func(t *testing.T) {
		// Set the expected behavior for the mock
		mockRepository.On("FindDisbursementByRequestID", mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockRepository.On("FindWalletForUpdateByUserCode", mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("string")).Return(mockWallet, nil).Once()
		mockRepository.On("Transaction").Return(sql.Tx{}, nil).Once()
		mockRepository.On("Rollback", mock.AnythingOfType("*sql.Tx")).Return(nil).Once()

		// Call your function for test
		_, err := services.Disbursement(mockInsufficientRequest)

		// Assert a result as expectations
		assert.Equal(t, err, errors.New(exception.InsufficientBalance))

		// Assert that the expectations for the mock were met
		mockRepository.AssertExpectations(t)
	})
	t.Run("DUPLICATE REQUEST", func(t *testing.T) {
		// Set the expected behavior for the mock
		mockRepository.On("FindDisbursementByRequestID", mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("string")).Return(mockDisbursement, nil).Once()
		//mockRepository.On("FindWalletForUpdateByUserCode", mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("string")).Return(mockWallet, nil).Once()
		mockRepository.On("Transaction").Return(sql.Tx{}, nil).Once()
		mockRepository.On("Rollback", mock.AnythingOfType("*sql.Tx")).Return(nil).Once()

		// Call your function for test
		_, err := services.Disbursement(mockInsufficientRequest)

		// Assert a result as expectations
		assert.Equal(t, err, errors.New(exception.DuplicateRequest))

		// Assert that the expectations for the mock were met
		mockRepository.AssertExpectations(t)
	})
	t.Run("SUCCESS", func(t *testing.T) {
		// Set the expected behavior for the mock
		mockRepository.On("FindDisbursementByRequestID", mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("string")).Return(nil, nil).Once()
		mockRepository.On("FindWalletForUpdateByUserCode", mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("string")).Return(mockWallet, nil).Once()
		mockRepository.On("UpdateWalletBalanceByUserCode", mock.AnythingOfType("*sql.Tx"), mock.AnythingOfType("int64"), mock.AnythingOfType("string")).Return(nil).Once()
		mockRepository.On(
			"InsertDisbursement",
			mock.AnythingOfType("*sql.Tx"),
			mock.AnythingOfType("*model.Disbursement"),
		).
			Return(mockDisbursement, nil).
			Once()
		mockRepository.On("Transaction").Return(sql.Tx{}, nil).Once()
		mockRepository.On("Rollback", mock.AnythingOfType("*sql.Tx")).Return(nil).Once()
		mockRepository.On("Commit", mock.AnythingOfType("*sql.Tx")).Return(nil).Once()

		// Call your function for test
		data, err := services.Disbursement(mockRequest)

		// Assert that the interaction with the mock was successful
		if err != nil {
			assert.NoError(t, err)
		} else {
			assert.NoError(t, nil)
		}

		// Assert a result as expectations
		if data != nil {
			assert.Equal(t, mockDisbursementResponse, data)
		}

		// Assert that the expectations for the mock were met
		mockRepository.AssertExpectations(t)
	})
}
