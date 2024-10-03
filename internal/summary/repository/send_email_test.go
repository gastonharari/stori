package repository

import (
	"context"
	"net/http"
	"testing"

	"stori/internal/summary/domain"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail_Success(t *testing.T) {
	tw := newTestWrapper(t)

	emailData := domain.EmailData{
		FromName:         "Sender",
		From:             "sender@example.com",
		ToName:           "Recipient",
		To:               "recipient@example.com",
		Subject:          "Test Subject",
		PlainTextContent: "This is a test email.",
		HTMLContent:      "<p>This is a test email.</p>",
	}
	from := mail.NewEmail(emailData.FromName, emailData.From)
	to := mail.NewEmail(emailData.ToName, emailData.To)
	message := mail.NewSingleEmail(from, emailData.Subject, to, emailData.PlainTextContent, emailData.HTMLContent)
	tw.mockSendGridClient.On("Send", message).Return(&rest.Response{StatusCode: http.StatusAccepted}, nil)

	err := tw.Repository.SendEmail(context.Background(), emailData)
	assert.NoError(t, err)
	tw.mockSendGridClient.AssertExpectations(t)
}

func TestSendEmail_ErrorSending(t *testing.T) {
	tw := newTestWrapper(t)

	emailData := domain.EmailData{
		FromName:         "Sender",
		From:             "sender@example.com",
		ToName:           "Recipient",
		To:               "recipient@example.com",
		Subject:          "Test Subject",
		PlainTextContent: "This is a test email.",
		HTMLContent:      "<p>This is a test email.</p>",
	}
	from := mail.NewEmail(emailData.FromName, emailData.From)
	to := mail.NewEmail(emailData.ToName, emailData.To)
	message := mail.NewSingleEmail(from, emailData.Subject, to, emailData.PlainTextContent, emailData.HTMLContent)
	tw.mockSendGridClient.On("Send", message).Return(nil, assert.AnError)

	err := tw.Repository.SendEmail(context.Background(), emailData)
	assert.Error(t, err)
	assert.ErrorIs(t, err, assert.AnError)
	tw.mockSendGridClient.AssertExpectations(t)
}

func TestSendEmail_InvalidResponseCode(t *testing.T) {
	tw := newTestWrapper(t)

	emailData := domain.EmailData{
		FromName:         "Sender",
		From:             "sender@example.com",
		ToName:           "Recipient",
		To:               "recipient@example.com",
		Subject:          "Test Subject",
		PlainTextContent: "This is a test email.",
		HTMLContent:      "<p>This is a test email.</p>",
	}
	from := mail.NewEmail(emailData.FromName, emailData.From)
	to := mail.NewEmail(emailData.ToName, emailData.To)
	message := mail.NewSingleEmail(from, emailData.Subject, to, emailData.PlainTextContent, emailData.HTMLContent)
	tw.mockSendGridClient.On("Send", message).Return(&rest.Response{StatusCode: http.StatusForbidden}, nil)

	err := tw.Repository.SendEmail(context.Background(), emailData)
	assert.Error(t, err)
	assert.ErrorIs(t, err, domain.ErrEmailInvalidResponseCode)
	tw.mockSendGridClient.AssertExpectations(t)
}
