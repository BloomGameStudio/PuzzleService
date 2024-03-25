package models

import (
	"gorm.io/gorm"
)
    
type Puzzle struct {
    gorm.Model
    Name string `json:"name"`
	Solution string `json:"solution"`
	Reward int  `json:"reward"`
    PosX int    `json:"posX"`
    PosY int    `json:"posY"`
    PosZ int    `json:"posZ"`
    Solved bool `json:"solved"`
}
