package sendemail

type UseCase struct {
	Service Service
}

func NewUseCase(service Service) UseCase {
	return UseCase{
		Service: service,
	}
}
