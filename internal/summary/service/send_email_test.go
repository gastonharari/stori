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
		},
	}

	expectedText := "Total balance is: 39.74\nNumber of transactions in July: 2\nAverage debit amount: -15.38\nAverage credit amount: 35.25\n"

	plainTextContent := generatePlainTextContent(summary)
	assert.Equal(t, expectedText, plainTextContent)
}

func TestGenerateHTMLContent(t *testing.T) {
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
		},
	}

	expectedHTML := "\n        <html>\n        <body style=\"background-color: #004034; font-family: Arial, sans-serif; margin: 0; padding: 0; color: #FFFFFF;\">\n            <div style=\"width: 100%; max-width: 600px; margin: 0 auto; padding: 20px; text-align: center; background-color: #FFFFFF; color: #004034; border-radius: 10px;\">\n                <img src=\"https://drive.google.com/uc?export=view&id=1R8xYbhWBji1k4V8ArTb667EKJyDJn99J\" alt=\"Stori Logo\" style=\"width: 150px; margin-bottom: 20px;\" />\n                <div style=\"padding: 20px;\">\n                    <h2 style=\"color: #00F08C;\">Monthly Summary</h2>\n                    <p><strong>Total balance is: 39.74</strong></p>\n                    <p><strong>Average debit amount:</strong> -15.38</p>\n                    <p><strong>Average credit amount:</strong> 35.25</p><p>Number of transactions in July: 2</p>\n                </div>\n            </div>\n        </body>\n        </html>"

	htmlContent := generateHTMLContent(summary)
	assert.Equal(t, expectedHTML, htmlContent)
}
