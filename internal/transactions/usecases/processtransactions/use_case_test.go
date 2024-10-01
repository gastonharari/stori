package processtransactions

import (
	"context"
	"stori/internal/transactions/usecases/processtransactions/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWrapper struct {
	uc                  UseCase
	ctx                 context.Context
	mockService         *mocks.Service
	mockCreateSummaryUC *mocks.CreateSummaryUC
	mockSendEmailUC     *mocks.SendEmailUC
}

func newTestWrapper(t *testing.T) testWrapper {
	t.Helper()
	mockService := mocks.NewService(t)
	mockCreateSummaryUC := mocks.NewCreateSummaryUC(t)
	mockSendEmailUC := mocks.NewSendEmailUC(t)

	uc := NewUseCase(mockService, mockCreateSummaryUC, mockSendEmailUC)

	return testWrapper{
		uc:                  uc,
		ctx:                 context.Background(),
		mockService:         mockService,
		mockCreateSummaryUC: mockCreateSummaryUC,
		mockSendEmailUC:     mockSendEmailUC,
	}
}

func TestNewUseCase(t *testing.T) {
	mockService := mocks.NewService(t)
	mockCreateSummaryUC := mocks.NewCreateSummaryUC(t)
	mockSendEmailUC := mocks.NewSendEmailUC(t)

	uc := NewUseCase(mockService, mockCreateSummaryUC, mockSendEmailUC)

	assert.NotNil(t, uc)
	assert.Equal(t, mockService, uc.Service)
	assert.Equal(t, mockCreateSummaryUC, uc.CreateSummaryUC)
	assert.Equal(t, mockSendEmailUC, uc.SendEmailUC)
}
