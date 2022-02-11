package model

import (
	"gorm.io/gorm"
	"time"
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
	LastLogin time.Time  `json:"lastLogin"`
	//Refer
	//2FA
}
