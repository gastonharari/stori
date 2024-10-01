package createsummary

import (
	"context"
	"testing"
)

type testWrapper struct {
	uc  UseCase
	ctx context.Context
}

func newTestWrapper(t *testing.T) testWrapper {
	t.Helper()
	return testWrapper{
		uc:  NewUseCase(),
		ctx: context.Background(),
	}
}
