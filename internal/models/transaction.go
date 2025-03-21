package models

import "time"

type Transaction struct {
	ID            int64     `bun:"id,pk,autoincrement" json:"id"`
	CustomerId    int64     `bun:"customer_id,notnull" json:"customer_id"`
	DeviceId      int64     `bun:"device_id,notnull" json:"device_id"`
	DeviceIssueId *int64    `bun:"device_issue_id" json:"device_issue_id,omitempty"`
	Note          *string   `bun:"note" json:"note,omitempty"` // Use pointer type to handle NULL values
	Status        string    `bun:"status,notnull,default:'Pending'" json:"status"`
	PickupAddress string    `bun:"pickup_address,notnull" json:"pickup_address"`
	FullAddress   *string   `bun:"full_address,notnull" json:"full_address"`
	CreatedAt     time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`
}

type TransactionWithDevice struct {
	ID               int       `bun:"id,pk,autoincrement" json:"id"`
	Status           string    `bun:"status,notnull,default:'Pending'" json:"status"`
	PickupAddress    string    `bun:"pickup_address,notnull" json:"pickup_address"`
	FullAddress      *string   `bun:"full_address,notnull" json:"full_address"`
	DeviceBrand      string    `bun:"brand,notnull" json:"brand"`
	DeviceModel      string    `bun:"model,notnull" json:"model"`
	IssueDescription string    `bun:"issue_description,notnull" json:"issue_description"`
	CreatedAt        time.Time `bun:"created_at,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt        time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updated_at"`
}
