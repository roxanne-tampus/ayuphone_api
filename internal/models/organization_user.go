package models

import "time"

type OrganizationUser struct {
	ID             int64     `bun:"id,pk,autoincrement"`
	UserID         int64     `bun:"user_id,notnull"`
	OrganizationID int64     `bun:"organization_id,notnull"`
	Role           string    `bun:"role,notnull,default:'technician'"` // 'admin', 'technician'
	CreatedAt      time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt      time.Time `bun:"updated_at,default:current_timestamp"`
}
