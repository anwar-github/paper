package test

import (
	"github.com/stretchr/testify/mock"
	"papper/service"
)

type MockRepository struct {
	mock.Mock
}

var (
	mockRepository = new(MockRepository)
	services       = service.NewService(mockRepository)
)
