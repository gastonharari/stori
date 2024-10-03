package service

import (
	"stori/internal/summary/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWrapper struct {
	service  Service
	mockRepo *mocks.Repository
}

func newTestWrapper(t *testing.T) testWrapper {
	t.Helper()
	mockRepo := mocks.NewRepository(t)
	fromEmail := "from@email.com"
	fromEmailName := "From Name"
	subject := "Subject"
	service := NewService(mockRepo, fromEmail, fromEmailName, subject)
	return testWrapper{
		service:  service,
		mockRepo: mockRepo,
	}
}

func TestNewService(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	fromEmail := "from@email.com"
	fromEmailName := "From Name"
	subject := "Subject"
	service := NewService(mockRepo, fromEmail, fromEmailName, subject)

	assert.NotNil(t, service)
	assert.Equal(t, mockRepo, service.Repository)
}
