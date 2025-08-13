package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email      string    `gorm:"type:text;unique;not null"`
	Username   string    `gorm:"type:varchar(20);unique;not null"`
	Firstname  string    `gorm:"type:varchar(15);not null"`
	Lastname   string    `gorm:"type:varchar(50);not null"`
	Password   string    `gorm:"type:text;not null"`
	IsAdmin    bool      `gorm:"not null;default:false"`
	IsVerified bool      `gorm:"not null;default:false"`
	IsBlocked  bool      `gorm:"not null;default:false"`
	CreatedAt  time.Time `gorm:"default:now()"`
	UpdatedAt  time.Time `gorm:"default:now()"`
}
type UserProfile struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	IsAdmin   bool
}
type RegisterReq struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Password  string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
