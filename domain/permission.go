package domain

import "time"

//Permission ...
type Permission struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Route     string    `json:"route"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
