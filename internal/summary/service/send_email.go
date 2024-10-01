package service

import (
	"context"
	"fmt"
	"stori/internal/summary/domain"
)

func (s Service) SendEmail(ctx context.Context, summary domain.Summary, userEmail string) error {
	htmlContent := generateHTMLContent(summary)
	plainTextContent := generatePlainTextContent(summary)
	emailData := domain.EmailData{
		From:             s.FromEmail,
		FromName:         s.FromEmailName,
		To:               userEmail,
		ToName:           userEmail,
		Subject:          s.Subject,
		PlainTextContent: plainTextContent,
		HTMLContent:      htmlContent,
	}

	err := s.Repository.SendEmail(ctx, emailData)
	if err != nil {
		return err
	}

	return nil
}

func generateHTMLContent(summary domain.Summary) string {
	html := fmt.Sprintf(`
        <html>
        <body style="background-color: #004034; font-family: Arial, sans-serif; margin: 0; padding: 0; color: #FFFFFF;">
            <div style="width: 100%%; max-width: 600px; margin: 0 auto; padding: 20px; text-align: center; background-color: #FFFFFF; color: #004034; border-radius: 10px;">
                <img src="https://drive.google.com/uc?export=view&id=1R8xYbhWBji1k4V8ArTb667EKJyDJn99J" alt="Stori Logo" style="width: 150px; margin-bottom: 20px;" />
                <div style="padding: 20px;">
                    <h2 style="color: #00F08C;">Monthly Summary</h2>
                    <p><strong>Total balance is: %.2f</strong></p>
                    <p><strong>Average debit amount:</strong> %.2f</p>
                    <p><strong>Average credit amount:</strong> %.2f</p>`, summary.TotalBalance, summary.AverageDebit, summary.AverageCredit)

	for month, monthSummary := range summary.MonthlySummary {
		html += fmt.Sprintf(`<p>Number of transactions in %s: %d</p>`, month, monthSummary.TransactionCount)
	}

	html += `
                </div>
            </div>
        </body>
        </html>`

	return html
}

func generatePlainTextContent(summary domain.Summary) string {
	text := fmt.Sprintf("Total balance is: %.2f\n", summary.TotalBalance)

	for month, monthSummary := range summary.MonthlySummary {
		text += fmt.Sprintf("Number of transactions in %s: %d\n", month, monthSummary.TransactionCount)
	}

	text += fmt.Sprintf("Average debit amount: %.2f\n", summary.AverageDebit)
	text += fmt.Sprintf("Average credit amount: %.2f\n", summary.AverageCredit)

	return text
}
