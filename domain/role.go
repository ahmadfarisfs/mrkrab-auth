package domain

import (
	"context"
	"time"
)

//Role ...
type Role struct {
	ID             uint64           `json:"id"`
	RoleName       string           `json:"role_name"`
	RolePermission []RolePermission `json:"role_permission"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

//RoleRepository Contract representation for role repository
type RoleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Role, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (r Role, err error)
	Update(ctx context.Context, ar *Role) (err error)
	Store(ctx context.Context, a *Role) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

//RoleUsecase usecase representation for role
type RoleUsecase interface {
	CreateRole(ctx context.Context, email, password string) (u User, token string, err error)
	AttachOrDetachRoleToPermission(ctx context.Context, roles *Role, permission *[]Permission) (u User, err error)
	AssignRoleToUser(ctx context.Context, User *[]User) (u User, err error)
}
