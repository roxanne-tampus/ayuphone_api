package models

import "time"

type Transaction struct {
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	CustomerId    int64     `bun:"customer_id,notnull" json:"customer_id"`
	DeviceId      int64     `bun:"device_id,notnull" json:"device_id"`
	DeviceIssueId *int64    `bun:"device_issue_id" json:"device_issue_id,omitempty"`
	Note          *string   `bun:"note" json:"note,omitempty"` // Use pointer type to handle NULL values
	Status        string    `bun:"status,notnull,default:'Pending'" json:"status"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`
}
