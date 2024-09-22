package accountTransaction

import "time"

type AccountTransactionFilter struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	FromCreatedAt *time.Time `json:"from_created_at"`
	ToCreatedAt   *time.Time `json:"to_created_at"`
}
