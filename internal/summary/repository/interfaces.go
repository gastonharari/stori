package repository

import (
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//go:generate mockery --name=SendGridClient --output=./mocks --structname=SendGridClient --filename=send_grid_client.go
type SendGridClient interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
}
