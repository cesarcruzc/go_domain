package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID        string     `json:"id" gorm:"type:char(36);primary_key;not null;unique_index"`
	Name      string     `json:"name" gorm:"type:varchar(50);not null"`
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (c *Course) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}

	return
}
