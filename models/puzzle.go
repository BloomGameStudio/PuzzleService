package models

import (
	"gorm.io/gorm"
)

type Puzzle struct {
    gorm.Model
    Name string
}