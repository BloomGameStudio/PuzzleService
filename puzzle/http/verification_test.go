package http

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BloomGameStudio/PuzzleService/database"
	publicmodels "github.com/BloomGameStudio/PuzzleService/publicModels"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup an in-memory database for testing
func setupTestDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %s", err)
	}

	// Auto migrate the database schema
	err = db.AutoMigrate(&publicmodels.Puzzle{})
	if err != nil {
		t.Fatalf("Failed to migrate database schema: %s", err)
	}

	// Seed the database with test data
	db.Create(&publicmodels.Puzzle{ID: 1, Data: "Correct!"})

	// Replace the global database instance with the in-memory one
	database.DB = db
	log.Println("In-memory database setup complete")
}

// Test verify route using correct solution
func TestVerifySolutionHandlerWithCorrectSolution(t *testing.T) {
	app := fiber.New()
	setupTestDB(t)

	app.Post("/verify", VerifySolutionHandler)

	body := bytes.NewBufferString(`{"ID":1,"Data":"Correct!"}`)
	req := httptest.NewRequest("POST", "/verify", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %s", resp.Status)
	} else {
		log.Println("TestVerifySolutionHandlerWithCorrectSolution passed")
	}
}

// Test verify route using incorrect solution for a given ID
func TestVerifySolutionHandlerWithIncorrectSolution(t *testing.T) {
	app := fiber.New()
	setupTestDB(t)

	app.Post("/verify", VerifySolutionHandler)

	body := bytes.NewBufferString(`{"ID":1,"Data":"Wrong!"}`)
	req := httptest.NewRequest("POST", "/verify", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %s", resp.Status)
	} else {
		log.Println("TestVerifySolutionHandlerWithIncorrectSolution passed")
	}
}

// Test verify route using invalid request
func TestVerifySolutionHandlerWithInvalidRequest(t *testing.T) {
	app := fiber.New()
	setupTestDB(t)

	app.Post("/verify", VerifySolutionHandler)

	body := bytes.NewBufferString(`{"ID":0}`)
	req := httptest.NewRequest("POST", "/verify", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status BadRequest, got %s", resp.Status)
	} else {
		log.Println("TestVerifySolutionHandlerWithInvalidRequest passed")
	}
}

// Test verify route using no solution found
func TestVerifySolutionHandlerWithNoSolutionFound(t *testing.T) {
	app := fiber.New()
	setupTestDB(t)

	app.Post("/verify", VerifySolutionHandler)

	body := bytes.NewBufferString(`{"ID":999,"Data":"Correct!"}`)
	req := httptest.NewRequest("POST", "/verify", body)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Failed to send request: %s", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status NotFound, got %s", resp.Status)
	} else {
		log.Println("TestVerifySolutionHandlerWithNoSolutionFound passed")
	}
}
