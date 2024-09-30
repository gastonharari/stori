package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddMultipleCreditsAndDebitsToDifferentMonths(t *testing.T) {
	initialSummary := Summary{
		TotalBalance:   0,
		MonthlySummary: make(map[string]MonthSummary),
	}

	transactions := []Transaction{
		{ID: "0", Date: time.Date(0, 7, 15, 0, 0, 0, 0, time.UTC), Amount: 60.50},
		{ID: "1", Date: time.Date(0, 7, 28, 0, 0, 0, 0, time.UTC), Amount: -10.3},
		{ID: "2", Date: time.Date(0, 8, 2, 0, 0, 0, 0, time.UTC), Amount: -20.46},
		{ID: "3", Date: time.Date(0, 8, 13, 0, 0, 0, 0, time.UTC), Amount: 10},
	}

	expectedSummary := Summary{
		TotalBalance:     39.74,
		AverageCredit:    35.25,
		AverageDebit:     -15.38,
		TotalDebitCount:  2,
		TotalCreditCount: 2,
		MonthlySummary: map[string]MonthSummary{
			"July": {
				Month:            "July",
				AverageCredit:    60.5,
				AverageDebit:     -10.3,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
			"August": {
				Month:            "August",
				AverageCredit:    10,
				AverageDebit:     -20.46,
				CountCredit:      1,
				CountDebit:       1,
				TransactionCount: 2,
			},
		},
	}

	for _, transaction := range transactions {
		initialSummary.AddNewTransaction(transaction)
	}

	assert.Equal(t, expectedSummary, initialSummary)
}
