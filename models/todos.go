package models

import (
	"time"
)

type List struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(100)"`
	Description string    `json:"description" gorm:"type:varchar(1000)"`
	Sublists    []Sublist `json:"sublists" gorm:"foreignKey:ListID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Sublist struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(100)"`
	Description string    `json:"description" gorm:"type:varchar(1000)"`
	ListID      int       `json:"list_id"`
	List        List      `json:"-" gorm:"foreignKey:ListID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
