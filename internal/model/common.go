package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Status string `json:"Status" gorm:"not null"`
}

type Word struct {
	gorm.Model
	Label string `json:"label" gorm:"unique"`
	Data  string `json:"data" gorm:"not null"`
}
