package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"papper/exception"
	"testing"
)

func TestWalletValidationService(t *testing.T) {
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
}
