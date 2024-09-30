package createsummary

type UseCase struct {
	service Service
}

func NewUseCase(service Service) UseCase {
	return UseCase{
		service: service,
	}
}
