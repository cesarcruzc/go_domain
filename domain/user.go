package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `json:"id" gorm:"type:char(36);primary_key;not null;unique_index"`
	FirstName string         `json:"firstName" gorm:"type:varchar(50);not null"`
	LastName  string         `json:"lastName" gorm:"type:varchar(50);not null"`
	Email     string         `json:"email" gorm:"type:varchar(100);not null;unique_index"`
	Phone     string         `json:"phone" gorm:"type:varchar(20);not null;unique_index"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	return
}
