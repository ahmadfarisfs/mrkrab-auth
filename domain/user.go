package domain

import (
	"context"
	"time"
)

//User ...
type User struct {
	ID        uint64    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	RoleID    uint64    `json:"role_id"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//UserRepository contract for repository user's
type UserRepository interface {
	Fetch(ctx context.Context, cursor string, num uint64) (res []User, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (u User, err error)
	GetByEmail(ctx context.Context, title string) (u User, err error)
	GetByRole(ctx context.Context, title string) (u User, err error)
	GetByPermission(ctx context.Context, title string) (u User, err error)
	Update(ctx context.Context, ar *User) (err error)
	Store(ctx context.Context, a *User) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

//UserUsecase All representation of user usecase
type UserUsecase interface {
	Fetch(ctx context.Context, cursor uint64, num uint64) (u []User, err error)
	Login(ctx context.Context, user *User) (u User, token string, err error)
	Register(ctx context.Context, user *User) (u User, err error)
	UpdateUser(ctx context.Context, user *User) (u User, err error)
	GetByID(ctx context.Context, id uint64) (u User, err error)
	AssignRole(ctx context.Context, role Role) (u User, err error)
	GetByRole(ctx context.Context, role Role) (u []User, err error)
	GetByPermission(ctx context.Context, permission Permission) (u []User, err error)
}
