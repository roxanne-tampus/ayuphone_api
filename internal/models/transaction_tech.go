package models

import (
	"time"
)

type TransactionTechnician struct {
	ID            int64      `bun:",pk,autoincrement" json:"id"`
	TransactionID int64      `bun:"transaction_id,notnull" json:"transaction_id"` // Foreign key to transactions table
	TechnicianID  int64      `bun:"technician_id,notnull" json:"technician_id"`   // Foreign key to users table
	AssignedAt    time.Time  `bun:"assigned_at,notnull,default:current_timestamp" json:"assigned_at"`
	UnassignedAt  *time.Time `bun:"unassigned_at" json:"unassignedAt,omitempty"` // Nullable for current assignment
}
