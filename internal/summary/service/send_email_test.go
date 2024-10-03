package service

import (
	"context"
	"stori/internal/summary/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail_Success(t *testing.T) {
	tw := newTestWrapper(t)

	summary := domain.Summary{
		TotalBalance:     39.74,
		AverageCredit:    35.25,
		AverageDebit:     -15.38,
		TotalDebitCount:  2,
		TotalCreditCount: 2,
		MonthlySummary: map[string]domain.MonthSummary{
			"July": {
				Month:            "July",
				AverageCredit:    60.5,
				AverageDebit:     -10.3,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
			"August": {
				Month:            "August",
				AverageCredit:    10,
				AverageDebit:     -20.46,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
		},
	}
	userEmail := "test@email.com"

	tw.mockRepo.On("SendEmail", context.Background(), domain.EmailData{
		From:             tw.service.FromEmail,
		FromName:         tw.service.FromEmailName,
		To:               userEmail,
		ToName:           userEmail,
		Subject:          tw.service.Subject,
		PlainTextContent: generatePlainTextContent(summary),
		HTMLContent:      generateHTMLContent(summary),
	}).Return(nil)

	err := tw.service.SendEmail(context.Background(), summary, userEmail)
	assert.NoError(t, err)
}

func TestSendEmail_Error(t *testing.T) {
	tw := newTestWrapper(t)

	summary := domain.Summary{
		TotalBalance:     39.74,
		AverageCredit:    35.25,
		AverageDebit:     -15.38,
		TotalDebitCount:  2,
		TotalCreditCount: 2,
		MonthlySummary: map[string]domain.MonthSummary{
			"July": {
				Month:            "July",
				AverageCredit:    60.5,
				AverageDebit:     -10.3,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
			"August": {
				Month:            "August",
				AverageCredit:    10,
				AverageDebit:     -20.46,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
		},
	}
	userEmail := "test@email.com"

	tw.mockRepo.On("SendEmail", context.Background(), domain.EmailData{
		From:             tw.service.FromEmail,
		FromName:         tw.service.FromEmailName,
		To:               userEmail,
		ToName:           userEmail,
		Subject:          tw.service.Subject,
		PlainTextContent: generatePlainTextContent(summary),
		HTMLContent:      generateHTMLContent(summary),
	}).Return(assert.AnError)

	err := tw.service.SendEmail(context.Background(), summary, userEmail)
	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
}

func TestGeneratePlainTextContent(t *testing.T) {
	summary := domain.Summary{
		TotalBalance:     39.74,
		AverageCredit:    35.25,
		AverageDebit:     -15.38,
		TotalDebitCount:  2,
		TotalCreditCount: 2,
		MonthlySummary: map[string]domain.MonthSummary{
			"July": {
				Month:            "July",
				AverageCredit:    60.5,
				AverageDebit:     -10.3,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
			"August": {
				Month:            "August",
				AverageCredit:    10,
				AverageDebit:     -20.46,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
		},
	}

	expectedText := `Total balance is: 39.74
Number of transactions in July: 2
Number of transactions in August: 2
Average debit amount: -15.38
Average credit amount: 35.25
`

	plainTextContent := generatePlainTextContent(summary)
	assert.Equal(t, expectedText, plainTextContent)
}
