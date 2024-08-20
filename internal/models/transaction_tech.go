package models

import (
	"time"
)

type TransactionTechnician struct {
	ID            int64      `bun:",pk,autoincrement"`
	TransactionID int64      `bun:"transaction_id,notnull"` // Foreign key to transactions table
	TechnicianID  int64      `bun:"technician_id,notnull"`  // Foreign key to users table
	AssignedAt    time.Time  `bun:"assigned_at,notnull,default:current_timestamp"`
	UnassignedAt  *time.Time `bun:"unassigned_at"` // Nullable for current assignment
}
