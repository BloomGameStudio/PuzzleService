package handlers

import "gorm.io/gorm"

// Interface that takes a pointed gorm.DB and returns an error
type HandlerInterface interface {
    GetPuzzles(db *gorm.DB) error
    CreatePuzzle(db *gorm.DB) error
    UpdatePuzzle(db *gorm.DB) error
    DeletePuzzle(db *gorm.DB) error
}

// Puzzle struct that contains the name and id of a puzzle
type Puzzle struct {
    Name string
    Id int
}

// Get all puzzles from the DB
func (p *Puzzle) GetPuzzles(db *gorm.DB) error {
    result := db.Find(p)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// Create a puzzle in the DB
func (p *Puzzle) CreatePuzzle(db *gorm.DB) error {
    result := db.Create(p)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// Update a puzzle in the DB
func (p *Puzzle) UpdatePuzzle(db *gorm.DB) error {
    result := db.Save(p)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

// Delete a puzzle in the DB
func (p *Puzzle) DeletePuzzle(db *gorm.DB) error {
    result := db.Delete(p)
    if result.Error != nil {
        return result.Error
    }
    return nil
}