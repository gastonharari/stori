package processfiletransactions

import (
	"flag"
	"net/http"
	"os"
	"stori/cmd/processfiletransactions/dtos"
	"testing"

	"github.com/sendgrid/rest"
	"github.com/stretchr/testify/assert"
)

func TestHandle_Success(t *testing.T) {
	tw := newTestWrapper(t)
	csvContent := [][]string{
		{"Id", "Date", "Transaction"},
		{"0", "7/15", "60.50"},
		{"1", "7/28", "-10.3"},
	}
	filename := mockCSV(t, csvContent)

	os.Args = []string{"cmd", "--file=" + filename, "--email=test@example.com"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	message := mockEmail("test@example.com")
	mockResponse := rest.Response{StatusCode: http.StatusAccepted}
	tw.mockSendGridClient.On("Send", message).Return(&mockResponse, nil)

	err := tw.handler.Handle()
	assert.Nil(t, err)
}

func TestHandle_ErrorNoFile(t *testing.T) {
	tw := newTestWrapper(t)
	os.Args = []string{"cmd", "--email=test@example.com"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	err := tw.handler.Handle()
	assert.Error(t, err)
	assert.Equal(t, dtos.ErrMissingFile, err)
}

func TestHandle_ErrorNoEmail(t *testing.T) {
	tw := newTestWrapper(t)
	os.Args = []string{"cmd", "--file=test.csv"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	err := tw.handler.Handle()
	assert.Error(t, err)
	assert.Equal(t, dtos.ErrMissingEmail, err)
}

func TestHandle_ErrorSendEmail(t *testing.T) {
	tw := newTestWrapper(t)
	csvContent := [][]string{
		{"Id", "Date", "Transaction"},
		{"0", "7/15", "60.50"},
		{"1", "7/28", "-10.3"},
	}
	filename := mockCSV(t, csvContent)

	os.Args = []string{"cmd", "--file=" + filename, "--email=test@example.com"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	message := mockEmail("test@example.com")
	mockResponse := rest.Response{}
	tw.mockSendGridClient.On("Send", message).Return(&mockResponse, assert.AnError)

	err := tw.handler.Handle()
	assert.Error(t, err)
	assert.Equal(t, assert.AnError, err)
}
