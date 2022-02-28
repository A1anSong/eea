package model

import (
	"time"

	"gorm.io/gorm"
)

type AuthLevel int

var (
	AuthNone AuthLevel = 0
	AuthKYC  AuthLevel = 1
)

type User struct {
	*gorm.Model `json:"-"`
	//UserName  string
	FirstName string     `json:"firstName,omitempty"`
	LastName  string     `json:"lastName,omitempty"`
	Email     string     `json:"email,omitempty"`
	Password  string     `json:"password,omitempty"`
	Role      UserRole   `json:"role,omitempty"`
	Status    UserStatue `json:"status,omitempty"`
	AuthLevel AuthLevel  `json:"auth_level,omitempty"`
	LastLogin time.Time  `json:"lastLogin"`
	//Refer
	//2FA
}
