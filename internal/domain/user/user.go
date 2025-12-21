package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Role      Role      `json:"role"`
}

type Role int

const (
	RoleSuperAdmin = iota
	RoleAdmin
	RoleUser
	RoleGuest
)

func (r Role) String() string {
	switch r {
	case RoleUser:
		return "user"
	case RoleAdmin:
		return "admin"
	case RoleGuest:
		return "guest"
	default:
		return "unknown"
	}
}
