package service

type Service struct {
	Repository    Repository
	FromEmail     string
	FromEmailName string
	Subject       string
}

func NewService(repository Repository, fromEmail string, fromEmailName string, subject string) Service {
	return Service{
		Repository:    repository,
		FromEmail:     fromEmail,
		FromEmailName: fromEmailName,
		Subject:       subject,
	}
}
