package sendemail

import (
	"context"
	"fmt"
	"stori/internal/summary/domain"
)

func (uc UseCase) Exec(ctx context.Context, summary domain.Summary) error {
	fmt.Println("Sending email with summary:")
	fmt.Println("Average Credit ", summary.AverageCredit)
	fmt.Println("Average Debit ", summary.AverageDebit)
	fmt.Println("Total Credit ", summary.TotalCreditCount)
	fmt.Println("Total Debit ", summary.TotalDebitCount)
	fmt.Println("Total Balance ", summary.TotalBalance)
	fmt.Println("----------------------------------------------------------")
	for _, v := range summary.MonthlySummary {
		fmt.Println("Month: ", v.Month)
		fmt.Println("Average Credit ", v.AverageCredit)
		fmt.Println("Average Debit ", v.AverageDebit)
		fmt.Println("Count Credit ", v.CountCredit)
		fmt.Println("Count Debit ", v.CountDebit)
		fmt.Println("Total Transactions ", v.TransactionCount)
		fmt.Println("----------------------------------------------------------")
	}
	return nil
}
