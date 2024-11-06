package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type List struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"type:varchar(100)" validate:"required,max=100,alphanum"`
	Description string    `json:"description" gorm:"type:varchar(1000)" validate:"required,max=1000"`
	Sublists    []Sublist `json:"sublists" gorm:"foreignKey:ListId;references:ID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	File        []string  `json:"file" validate:"required,dive,extension=txt|pdf"`
}

type Sublist struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"type:varchar(100)" validate:"required,max=100,alphanum"`
	Description string    `json:"description" gorm:"type:varchar(1000)" validate:"required,max=1000"`
	ListID      int       `json:"list_id"`
	List        List      `json:"list" gorm:"foreignKey:ListID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	File        []string  `json:"file" validate:"required,dive,extension=txt|pdf"`
}

func (l *List) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()
	return validate.Struct(l)
}

func (s *Sublist) BeforeCreate(tx *gorm.DB) error {
	validate := validator.New()
	return validate.Struct(s)
}
