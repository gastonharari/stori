package domain

const MonthFormat = "January"

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

type Month struct {
	Month string
}
