package repository

type Repository struct {
	EmailClient SendGridClient
}

func NewRepository(emailClient SendGridClient) Repository {
	return Repository{
		EmailClient: emailClient,
	}
}
