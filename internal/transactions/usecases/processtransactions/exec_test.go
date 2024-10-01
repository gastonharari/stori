package processtransactions

import (
	"errors"
	"stori/internal/summary/domain"
	dmntransactions "stori/internal/transactions/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExec_Success(t *testing.T) {
	tw := newTestWrapper(t)
	path := "testdata/transactions.csv"
	userEmail := "test@email.com"

	mockTransactions := []dmntransactions.Transaction{
		{ID: "0", Date: time.Date(0, 7, 15, 0, 0, 0, 0, time.UTC), Amount: 60.50},
		{ID: "1", Date: time.Date(0, 7, 28, 0, 0, 0, 0, time.UTC), Amount: -10.3},
	}
	mockSummary := domain.Summary{
		TotalBalance:     50.20,
		AverageCredit:    60.50,
		AverageDebit:     -10.3,
		TotalDebitCount:  1,
		TotalCreditCount: 1,
		MonthlySummary: map[string]domain.MonthSummary{
			"July": {
				Month:            "July",
				AverageCredit:    60.5,
				AverageDebit:     -10.3,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
		},
	}

	tw.mockService.On("ReadFile", tw.ctx, path).Return(mockTransactions, nil)
	tw.mockCreateSummaryUC.On("Exec", tw.ctx, mockTransactions).Return(mockSummary)
	tw.mockSendEmailUC.On("Exec", tw.ctx, mockSummary, userEmail).Return(nil)

	err := tw.uc.Exec(tw.ctx, path, userEmail)

	assert.NoError(t, err)
	tw.mockService.AssertExpectations(t)
	tw.mockCreateSummaryUC.AssertExpectations(t)
	tw.mockSendEmailUC.AssertExpectations(t)
}

func TestExec_ReadFileError(t *testing.T) {
	tw := newTestWrapper(t)
	path := "testdata/transactions.csv"
	userEmail := "test@email.com"

	expectedError := errors.New("read file error")

	tw.mockService.On("ReadFile", tw.ctx, path).Return(nil, expectedError)

	err := tw.uc.Exec(tw.ctx, path, userEmail)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	tw.mockService.AssertExpectations(t)
}

func TestExec_SendEmailError(t *testing.T) {
	tw := newTestWrapper(t)
	path := "testdata/transactions.csv"
	userEmail := "test@email.com"

	transactions := []dmntransactions.Transaction{
		{ID: "0", Date: time.Date(2023, 7, 15, 0, 0, 0, 0, time.UTC), Amount: 60.50},
		{ID: "1", Date: time.Date(2023, 7, 28, 0, 0, 0, 0, time.UTC), Amount: -10.3},
	}
	expectedSummary := domain.Summary{
		TotalBalance:     50.20,
		AverageCredit:    60.50,
		AverageDebit:     -10.3,
		TotalDebitCount:  1,
		TotalCreditCount: 1,
		MonthlySummary: map[string]domain.MonthSummary{
			"July": {
				Month:            "July",
				AverageCredit:    60.5,
				AverageDebit:     -10.3,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
		},
	}
	expectedError := errors.New("send email error")

	tw.mockService.On("ReadFile", tw.ctx, path).Return(transactions, nil)
	tw.mockCreateSummaryUC.On("Exec", tw.ctx, transactions).Return(expectedSummary)
	tw.mockSendEmailUC.On("Exec", tw.ctx, expectedSummary, userEmail).Return(expectedError)

	err := tw.uc.Exec(tw.ctx, path, userEmail)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	tw.mockService.AssertExpectations(t)
	tw.mockCreateSummaryUC.AssertExpectations(t)
	tw.mockSendEmailUC.AssertExpectations(t)
}
