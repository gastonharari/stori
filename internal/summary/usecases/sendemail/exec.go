package sendemail

import (
	"context"
	"stori/internal/summary/domain"
)

func (uc UseCase) Exec(ctx context.Context, summary domain.Summary, userEmail string) error {
	return uc.Service.SendEmail(ctx, summary, userEmail)
}
