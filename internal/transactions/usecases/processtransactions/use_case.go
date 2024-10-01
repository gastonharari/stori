package processtransactions

type UseCase struct {
	Service         Service
	CreateSummaryUC CreateSummaryUC
	SendEmailUC     SendEmailUC
}

func NewUseCase(service Service, createSummaryUC CreateSummaryUC, sendEmailUC SendEmailUC) UseCase {
	return UseCase{
		Service:         service,
		CreateSummaryUC: createSummaryUC,
		SendEmailUC:     sendEmailUC,
	}
}
