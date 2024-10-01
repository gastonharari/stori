package processfiletransactions

import (
	"log"
	"os"
	rposummary "stori/internal/summary/repository"
	srvsummary "stori/internal/summary/service"
	"stori/internal/summary/usecases/createsummary"
	"stori/internal/summary/usecases/sendemail"
	rpotransactions "stori/internal/transactions/repository"
	srvtransactions "stori/internal/transactions/service"
	"stori/internal/transactions/usecases/processtransactions"

	"github.com/sendgrid/sendgrid-go"
)

type handler struct {
	UC ProcesstransactionsUC
}

func NewHandler() (handler, error) {
	repoTransactions := rpotransactions.NewRepository()
	serviceTransactions := srvtransactions.NewService(repoTransactions)
	createSummaryUC := createsummary.NewUseCase()

	sendGridClient := sendgrid.NewSendClient(getSendGridApiKey())
	repoSummary := rposummary.NewRepository(sendGridClient)
	serviceSummary := srvsummary.NewService(repoSummary, getFromEmail(), getFromEmailName(), getSubject())
	sendEmailUC := sendemail.NewUseCase(serviceSummary)

	useCase := processtransactions.NewUseCase(serviceTransactions, createSummaryUC, sendEmailUC)
	return handler{
		UC: useCase,
	}, nil
}

func getSendGridApiKey() string {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		log.Fatal("SENDGRID_API_KEY is not set")
	}
	return apiKey
}

func getFromEmail() string {
	fromEmail := os.Getenv("EMAIL_FROM")
	if fromEmail == "" {
		log.Fatal("EMAIL_FROM is not set")
	}
	return fromEmail
}

func getFromEmailName() string {
	fromEmailName := os.Getenv("EMAIL_FROM_NAME")
	if fromEmailName == "" {
		log.Fatal("EMAIL_FROM_NAME is not set")
	}
	return fromEmailName
}

func getSubject() string {
	subject := os.Getenv("EMAIL_SUBJECT")
	if subject == "" {
		log.Fatal("EMAIL_SUBJECT is not set")
	}
	return subject
}
