package processfiletransactions

import (
	"stori/internal/summary/usecases/createsummary"
	"stori/internal/summary/usecases/sendemail"
	rpotransactions "stori/internal/transactions/repository"
	srvtransactions "stori/internal/transactions/service"
	"stori/internal/transactions/usecases/processtransactions"
)

type handler struct {
	UC ProcesstransactionsUC
}

func NewHandler() (handler, error) {
	repoTransactions := rpotransactions.NewRepository()
	serviceTransactions := srvtransactions.NewService(repoTransactions)
	createSummaryUC := createsummary.NewUseCase()
	sendEmailUC := sendemail.NewUseCase()

	useCase := processtransactions.NewUseCase(serviceTransactions, createSummaryUC, sendEmailUC)
	return handler{
		UC: useCase,
	}, nil
}
