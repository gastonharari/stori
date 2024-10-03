package processfiletransactions

import (
	"encoding/csv"
	"os"
	"stori/cmd/processfiletransactions/mocks"
	rposummary "stori/internal/summary/repository"
	srvsummary "stori/internal/summary/service"
	"stori/internal/summary/usecases/createsummary"
	"stori/internal/summary/usecases/sendemail"
	rpotransactions "stori/internal/transactions/repository"
	srvtransactions "stori/internal/transactions/service"
	"stori/internal/transactions/usecases/processtransactions"
	"testing"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type testWrapper struct {
	handler            handler
	mockSendGridClient *mocks.SendGridClient
}

func newTestWrapper(t *testing.T) testWrapper {
	t.Helper()
	t.Setenv("SENDGRID_API_KEY", "test")
	t.Setenv("EMAIL_FROM", "test@email.com")
	t.Setenv("EMAIL_FROM_NAME", "Test")
	t.Setenv("EMAIL_SUBJECT", "Test")
	repoTransactions := rpotransactions.NewRepository()
	serviceTransactions := srvtransactions.NewService(repoTransactions)
	createSummaryUC := createsummary.NewUseCase()

	sendGridClient := mocks.NewSendGridClient(t)
	repoSummary := rposummary.NewRepository(sendGridClient)
	serviceSummary := srvsummary.NewService(repoSummary, getFromEmail(), getFromEmailName(), getSubject())
	sendEmailUC := sendemail.NewUseCase(serviceSummary)

	useCase := processtransactions.NewUseCase(serviceTransactions, createSummaryUC, sendEmailUC)

	return testWrapper{
		handler:            handler{UC: useCase},
		mockSendGridClient: sendGridClient,
	}
}

func mockCSV(t *testing.T, content [][]string) string {
	file, err := os.CreateTemp("", "testfile-*.csv")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	writer.Flush()

	return file.Name()
}

func mockEmail(toEmail string) *mail.SGMailV3 {
	fromEmail := os.Getenv("EMAIL_FROM")
	fromName := os.Getenv("EMAIL_FROM_NAME")
	subject := os.Getenv("EMAIL_SUBJECT")
	from := mail.NewEmail(fromName, fromEmail)
	to := mail.NewEmail(toEmail, toEmail)
	plainText := mockPlanText()
	htmlContent := mockHTMLContent()
	message := mail.NewSingleEmail(from, subject, to, plainText, htmlContent)
	return message
}

func mockHTMLContent() string {
	return "\n        <html>\n        <body style=\"background-color: #004034; font-family: Arial, sans-serif; margin: 0; padding: 0; color: #FFFFFF;\">\n            <div style=\"width: 100%; max-width: 600px; margin: 0 auto; padding: 20px; text-align: center; background-color: #FFFFFF; color: #004034; border-radius: 10px;\">\n                <img src=\"https://drive.google.com/uc?export=view&id=1R8xYbhWBji1k4V8ArTb667EKJyDJn99J\" alt=\"Stori Logo\" style=\"width: 150px; margin-bottom: 20px;\" />\n                <div style=\"padding: 20px;\">\n                    <h2 style=\"color: #00F08C;\">Monthly Summary</h2>\n                    <p><strong>Total balance is: 50.20</strong></p>\n                    <p><strong>Average debit amount:</strong> -10.30</p>\n                    <p><strong>Average credit amount:</strong> 60.50</p><p>Number of transactions in July: 2</p>\n                </div>\n            </div>\n        </body>\n        </html>"
}

func mockPlanText() string {
	return "Total balance is: 50.20\nNumber of transactions in July: 2\nAverage debit amount: -10.30\nAverage credit amount: 60.50\n"
}
