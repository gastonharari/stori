package sendemail

import (
	"stori/internal/summary/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecSendEmail_Success(t *testing.T) {
	tw := newTestWrapper(t)
	summary := domain.Summary{
		TotalBalance:    50.20,
		AverageCredit:   60.50,
		AverageDebit:    -10.3,
		TotalDebitCount: 1,
	}
	email := "test@email.com"

	tw.mockService.On("SendEmail", tw.ctx, summary, email).Return(nil)

	err := tw.uc.Exec(tw.ctx, summary, email)
	assert.NoError(t, err)
}

func TestExecSendEmail_SendEmailError(t *testing.T) {
	tw := newTestWrapper(t)
	summary := domain.Summary{
		TotalBalance:    50.20,
		AverageCredit:   60.50,
		AverageDebit:    -10.3,
		TotalDebitCount: 1,
	}
	email := "test@email.com"

	tw.mockService.On("SendEmail", tw.ctx, summary, email).Return(assert.AnError)

	err := tw.uc.Exec(tw.ctx, summary, email)
	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
}
