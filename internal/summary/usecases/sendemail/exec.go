package sendemail

import (
	"context"
	"fmt"
	"stori/internal/summary/domain"
)

func (uc UseCase) Exec(ctx context.Context, summary domain.Summary) error {
	fmt.Println("Sending email with summary:", summary)
	return nil
}
