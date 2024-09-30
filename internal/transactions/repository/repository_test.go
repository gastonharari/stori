package repository

import "context"

type testWrapper struct {
	ctx        context.Context
	repository Repository
}

func newTestWrapper() testWrapper {
	return testWrapper{
		ctx:        context.Background(),
		repository: NewRepository(),
	}
}
