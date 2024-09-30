package domain

import "time"

type Transaction struct {
	ID     string
	Date   time.Time
	Amount float64
	UserID string
}
