package models

import "time"

type User struct {
	ID          int64     `bun:"id,pk,autoincrement"`
	FirstName   string    `bun:"first_name"`
	LastName    string    `bun:"last_name"`
	Email       string    `bun:"email,unique"`
	PhoneNumber string    `bun:"phone_number,unique"`
	Password    string    `bun:"password,notnull"`
	RoleID      int       `bun:"role_id,notnull,default:3"` // 'superadmin', 'admin', 'customer', 'technician'
	CreatedAt   time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,default:current_timestamp"`
}
