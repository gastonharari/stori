package domain

const monthFormat = "January"

type Summary struct {
	TotalBalance     float64
	AverageDebit     float64
	AverageCredit    float64
	MonthlySummary   map[string]MonthSummary
	TotalDebitCount  int
	TotalCreditCount int
}

type MonthSummary struct {
	Month            string
	AverageDebit     float64
	AverageCredit    float64
	CountDebit       int
	CountCredit      int
	TransactionCount int
}

func (s *Summary) AddNewTransaction(t Transaction) {
	monthKey := t.Date.Format(monthFormat)

	monthSummary, exists := s.MonthlySummary[monthKey]
	if !exists {
		monthSummary = MonthSummary{Month: monthKey}
	}

	monthSummary.TransactionCount++

	if t.Amount > 0 {
		monthSummary.CountCredit++
		monthSummary.AverageCredit = s.CalculateNewAverage(monthSummary.AverageCredit, monthSummary.CountCredit, t.Amount)

		s.TotalCreditCount++
		s.AverageCredit = s.CalculateNewAverage(s.AverageCredit, s.TotalCreditCount, t.Amount)
	} else {
		monthSummary.CountDebit++
		monthSummary.AverageDebit = s.CalculateNewAverage(monthSummary.AverageDebit, monthSummary.CountDebit, t.Amount)

		s.TotalDebitCount++
		s.AverageDebit = s.CalculateNewAverage(s.AverageDebit, s.TotalDebitCount, t.Amount)
	}

	s.TotalBalance += t.Amount
	s.MonthlySummary[monthKey] = monthSummary
}

func (s Summary) CalculateNewAverage(currentAverage float64, currentCount int, newAmount float64) float64 {
	return ((currentAverage * float64(currentCount-1)) + newAmount) / float64(currentCount)
}
