package createsummary

import (
	"context"
	"stori/internal/transactions/domain"
)

func (uc UseCase) Exec(ctx context.Context, path string) (domain.Summary, error) {
	transactions, err := uc.service.ReadFile(ctx, path)
	if err != nil {
		return domain.Summary{}, err
	}

	summary := createSummaty(transactions)

	return summary, nil
}

func createSummaty(transactions []domain.Transaction) domain.Summary {
	summary := domain.Summary{
		MonthlySummary: make(map[string]domain.MonthSummary),
	}

	for _, t := range transactions {
		summary.AddNewTransaction(t)
	}

	return summary
}
