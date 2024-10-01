package domain

type EmailData struct {
	From             string
	FromName         string
	To               string
	ToName           string
	Subject          string
	HTMLContent      string
	PlainTextContent string
}
