package models

import "time"

type OrganizationUser struct {
	ID             int64     `bun:"id,pk,autoincrement" json:"-"`
	UserID         int64     `bun:"user_id,notnull" json:"user_id"`
	OrganizationID int64     `bun:"organization_id,notnull" json:"-"`
	Role           string    `bun:"role,notnull,default:'technician'" json:"role"` // 'admin', 'technician'
	InvitedBy      *int64    `bun:"invited_by" json:"invited_by,omitempty"`
	Status         string    `bun:"status,notnull,default:'pending'" json:"status"` //pending, approved, or declined
	CreatedAt      time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
}

type OrganizationUserWithNames struct {
	OrganizationUser OrganizationUser `json:"organization_user"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
}
