package repository

import (
	"context"
	"errors"
	"log"
	"net/http"
	"stori/internal/summary/domain"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (r Repository) SendEmail(ctx context.Context, email domain.EmailData) error {
	from := mail.NewEmail(email.FromName, email.From)
	to := mail.NewEmail(email.ToName, email.To)

	message := mail.NewSingleEmail(from, email.Subject, to, email.PlainTextContent, email.HTMLContent)

	response, err := r.EmailClient.Send(message)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusAccepted {
		log.Printf("Error sending email: %v", response.Body)
		log.Printf("Response code: %v", response.StatusCode)
		return errors.Join(domain.ErrEmailInvaLIdResponseCode)
	}

	return nil
}
