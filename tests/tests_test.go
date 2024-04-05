package tests

import (
	"testing"

	"github.com/BloomGameStudio/PuzzleService/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestPuzzleOperations tests the CreatePuzzle, GetPuzzles, UpdatePuzzle, and DeletePuzzle functions
// That are defined in the handlers/puzzleHandler.go interface

func TestPuzzleOperations(t *testing.T) {
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatalf("failed to open database: %v", err)
    }

    err = db.AutoMigrate(&handlers.Puzzle{})
    if err != nil {
        t.Fatalf("failed to migrate database: %v", err)
    }

    p := &handlers.Puzzle{Name: "Test", Id: 1}

    // Test CreatePuzzle
    err = p.CreatePuzzle(db)
    if err != nil {
        t.Errorf("failed to create puzzle: %v", err)
    }

    // Test GetPuzzles
    err = p.GetPuzzles(db)
    if err != nil {
        t.Errorf("failed to get puzzles: %v", err)
    }

    // Test UpdatePuzzle
    p.Name = "Updated"
    err = p.UpdatePuzzle(db)
    if err != nil {
        t.Errorf("failed to update puzzle: %v", err)
    }

    // Test DeletePuzzle
    err = p.DeletePuzzle(db)
    if err != nil {
        t.Errorf("failed to delete puzzle: %v", err)
    }
}