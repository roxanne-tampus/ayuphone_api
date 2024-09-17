package models

import "time"

type Organization struct {
	ID        int64     `bun:"id,pk,autoincrement" json:"id"`
	Name      string    `bun:"name,notnull" json:"name"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp" json:"updated_at"`
}
