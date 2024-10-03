package processtransactions

import (
	"context"
	"stori/internal/transactions/domain"
)

func (uc UseCase) Exec(ctx context.Context, userEmail string, transactions []domain.Transaction) error {
	err := uc.Service.Create(ctx, transactions)
	if err != nil {
		return err
	}

	summary := uc.CreateSummaryUC.Exec(ctx, transactions)

	err = uc.SendEmailUC.Exec(ctx, summary, userEmail)
	if err != nil {
		return err
	}

	return nil
}
