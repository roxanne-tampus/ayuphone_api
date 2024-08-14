package models

import "time"

type User struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	Email     string    `bun:"email,unique,notnull"`
	Password  string    `bun:"password,notnull"`
	Role      string    `bun:"role,notnull,default:'customer'"` // 'superadmin', 'admin', 'customer', 'technician'
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp"`
}
