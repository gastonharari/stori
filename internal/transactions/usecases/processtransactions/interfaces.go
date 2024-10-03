package processtransactions

import (
	"context"
	dmnsummary "stori/internal/summary/domain"
	dmntransactions "stori/internal/transactions/domain"
)

//go:generate mockery --name=Service --output=./mocks --structname=Service --filename=service.go
//go:generate mockery --name=CreateSummaryUC --output=./mocks --structname=CreateSummaryUC --filename=create_summary.go
//go:generate mockery --name=SendEmailUC --output=./mocks --structname=SendEmailUC --filename=send_email.go

type Service interface {
	Create(ctx context.Context, transactions []dmntransactions.Transaction) error
}

type CreateSummaryUC interface {
	Exec(ctx context.Context, transactions []dmntransactions.Transaction) dmnsummary.Summary
}

type SendEmailUC interface {
	Exec(ctx context.Context, summary dmnsummary.Summary, userEmail string) error
}
