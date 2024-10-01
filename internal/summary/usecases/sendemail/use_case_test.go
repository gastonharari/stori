package sendemail

import (
	"context"
	"stori/internal/summary/usecases/sendemail/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWrapper struct {
	uc          UseCase
	ctx         context.Context
	mockService *mocks.Service
}

func newTestWrapper(t *testing.T) testWrapper {
	t.Helper()
	mockService := mocks.NewService(t)

	return testWrapper{
		uc:          NewUseCase(mockService),
		ctx:         context.Background(),
		mockService: mockService,
	}
}

func TestNewUsecase(t *testing.T) {
	mockService := mocks.NewService(t)
	uc := NewUseCase(mockService)
	assert.NotNil(t, uc)
	assert.Equal(t, mockService, uc.Service)
}
