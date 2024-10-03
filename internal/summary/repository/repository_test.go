package repository

import (
	"stori/internal/summary/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWrapper struct {
	Repository         Repository
	mockSendGridClient *mocks.SendGridClient
}

func newTestWrapper(t *testing.T) testWrapper {
	t.Helper()
	mockSendGridClient := mocks.NewSendGridClient(t)
	repo := NewRepository(mockSendGridClient)
	return testWrapper{
		Repository:         repo,
		mockSendGridClient: mockSendGridClient,
	}
}

func TestNewRepositoty(t *testing.T) {
	mockSendGridClient := mocks.NewSendGridClient(t)

	repo := NewRepository(mockSendGridClient)

	assert.NotNil(t, repo)
	assert.Equal(t, mockSendGridClient, repo.EmailClient)
}
