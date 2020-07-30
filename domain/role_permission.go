package domain

import (
	"context"
	"time"
)

//Domain for repository no need for usecase because it not necessary for user to see this pivot tables
//RolePermission ...
type RolePermission struct {
	ID           uint64       `json:"id"`
	RoleID       uint64       `json:"role_id"`
	PermissionID uint64       `json:"permission_id"`
	Role         []Role       `json:"role"`
	Permission   []Permission `json:"permission"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

//RolePermissionRepository is repository contract for RolePermission's its a pivot tables
type RolePermissionRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []RolePermission, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (r Role, err error)
	Update(ctx context.Context, ar *RolePermission) (err error)
	Store(ctx context.Context, a *RolePermission) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
