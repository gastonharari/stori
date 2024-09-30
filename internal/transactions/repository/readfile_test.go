package repository

import (
	"encoding/csv"
	"os"
	"testing"
	"time"

	"stori/internal/transactions/domain"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	tw := newTestWrapper()
	tests := []struct {
		name        string
		content     [][]string
		expected    []domain.Transaction
		expectError error
	}{
		{
			name: "valid file",
			content: [][]string{
				{"0", "7/15", "60.50"},
				{"1", "7/28", "-10.3"},
				{"2", "8/2", "-20.46"},
				{"3", "8/13", "+10"},
			},
			expected: []domain.Transaction{
				{ID: "0", Date: time.Date(0, 7, 15, 0, 0, 0, 0, time.UTC), Amount: 60.50},
				{ID: "1", Date: time.Date(0, 7, 28, 0, 0, 0, 0, time.UTC), Amount: -10.3},
				{ID: "2", Date: time.Date(0, 8, 2, 0, 0, 0, 0, time.UTC), Amount: -20.46},
				{ID: "3", Date: time.Date(0, 8, 13, 0, 0, 0, 0, time.UTC), Amount: 10},
			},
			expectError: nil,
		},
		{
			name:        "file not found",
			content:     nil,
			expected:    nil,
			expectError: domain.ErrorFileNotFound,
		},
		{
			name: "invalid id",
			content: [][]string{
				{"", "12/28", "100.50"},
			},
			expected:    nil,
			expectError: domain.ErrorInvalidID,
		},
		{
			name: "invalid date format",
			content: [][]string{
				{"1", "invalid-date", "100.50"},
			},
			expected:    nil,
			expectError: domain.ErrorInvalidDate,
		},
		{
			name: "invalid amount format",
			content: [][]string{
				{"1", "12/28", "invalid-amount"},
			},
			expected:    nil,
			expectError: domain.ErrorInvalidAmount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filename string
			if tt.content != nil {
				filename = mockCSV(t, tt.content)
				defer os.Remove(filename)
			} else {
				filename = "nonexistentfile.csv"
			}

			result, err := tw.repository.ReadFile(tw.ctx, filename)

			if tt.expectError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func mockCSV(t *testing.T, content [][]string) string {
	file, err := os.CreateTemp("", "testfile-*.csv")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	writer.Flush()

	return file.Name()
}
