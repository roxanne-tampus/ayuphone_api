package models

import "time"

type Transaction struct {
	ID            int64     `bun:"id,pk,autoincrement"`
	CustomerId    int64     `bun:"customer_id,notnull"`
	DeviceId      int64     `bun:"device_id,notnull"`
	DeviceIssueId int64     `bun:"device_issue_id,notnull"`
	Note          *string   `bun:"note"` // Use pointer type to handle NULL values
	Status        string    `bun:"status,notnull,default:'Pending'"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp"`
}
