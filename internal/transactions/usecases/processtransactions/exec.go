package processtransactions

import (
	"context"
)

func (uc UseCase) Exec(ctx context.Context, path string) error {
	transactions, err := uc.Service.ReadFile(ctx, path)
	if err != nil {
		return err
	}

	summary := uc.CreateSummaryUC.Exec(ctx, transactions)

	err = uc.SendEmailUC.Exec(ctx, summary)
	if err != nil {
		return err
	}

	return nil
}
