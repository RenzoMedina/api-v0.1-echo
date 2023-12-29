package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Code        string `gorm:"typ:varchar(10); not null"`
	Name        string `gorm:"typ:varchar(50); not null"`
	Description string `gorm:"typ:text; not null"`
	Stock       string `gorm:"typ:varchar(10); not null"`
}
