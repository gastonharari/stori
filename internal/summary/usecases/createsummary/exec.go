package createsummary

import (
	"context"
	dmnsummary "stori/internal/summary/domain"
	dmntransactions "stori/internal/transactions/domain"
)

func (uc UseCase) Exec(ctx context.Context, transactions []dmntransactions.Transaction) dmnsummary.Summary {
	summary := dmnsummary.Summary{
		MonthlySummary: make(map[string]dmnsummary.MonthSummary),
	}

	for _, t := range transactions {
		addNewTransaction(&summary, t)
	}

	return summary
}

func addNewTransaction(s *dmnsummary.Summary, t dmntransactions.Transaction) {
	monthKey := t.Date.Format(dmnsummary.MonthFormat)

	monthSummary, exists := s.MonthlySummary[monthKey]
	if !exists {
		monthSummary = dmnsummary.MonthSummary{Month: monthKey}
	}

	monthSummary.TransactionCount++

	if t.Amount > 0 {
		monthSummary.CountCredit++
		monthSummary.AverageCredit = calculateNewAverage(monthSummary.AverageCredit, monthSummary.CountCredit, t.Amount)

		s.TotalCreditCount++
		s.AverageCredit = calculateNewAverage(s.AverageCredit, s.TotalCreditCount, t.Amount)
	} else {
		monthSummary.CountDebit++
		monthSummary.AverageDebit = calculateNewAverage(monthSummary.AverageDebit, monthSummary.CountDebit, t.Amount)

		s.TotalDebitCount++
		s.AverageDebit = calculateNewAverage(s.AverageDebit, s.TotalDebitCount, t.Amount)
	}

	s.TotalBalance += t.Amount
	s.MonthlySummary[monthKey] = monthSummary
}

func calculateNewAverage(currentAverage float64, currentCount int, newAmount float64) float64 {
	return ((currentAverage * float64(currentCount-1)) + newAmount) / float64(currentCount)
}
