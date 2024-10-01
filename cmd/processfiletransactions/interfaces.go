package processfiletransactions

import "context"

type ProcesstransactionsUC interface {
	Exec(ctx context.Context, path string) error
}
