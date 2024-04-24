package publicmodels

import "gorm.io/gorm"

type Puzzle struct {
	gorm.Model

	Name string
	Data string
}
