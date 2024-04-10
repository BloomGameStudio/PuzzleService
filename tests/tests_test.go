package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/BloomGameStudio/PuzzleService/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestPuzzleHandler is a test function that tests all the CRUD operations of the PuzzleHandler.
func TestPuzzleHandler(t *testing.T) {

    // Initialize a new Fiber app and PuzzleHandler.
    app := fiber.New()
    puzzleHandlers := handlers.NewPuzzleHandler()

    // Set up the routes for the PuzzleHandler.
    app.Get("/puzzles/:id", puzzleHandlers.GetPuzzles)
    app.Post("/puzzles", puzzleHandlers.CreatePuzzle)
    app.Put("/puzzles/:id", puzzleHandlers.UpdatePuzzle)
    app.Delete("/puzzles/:id", puzzleHandlers.DeletePuzzle)

    // Test CreatePuzzle by sending a POST request with a new puzzle.
    req := httptest.NewRequest(http.MethodPost, "/puzzles", strings.NewReader(`{"id":"1","name":"test","posX":1,"posY":2,"posZ":3,"solution":"test","solved":false}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    // Test GetPuzzles by sending a GET request for the puzzle we just created.
    req = httptest.NewRequest(http.MethodGet, "/puzzles/1", nil)
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    body, _ := io.ReadAll(resp.Body)
    assert.Contains(t, string(body), `"id":"1"`)

    // Test UpdatePuzzle by sending a PUT request with updated puzzle data.
    req = httptest.NewRequest(http.MethodPut, "/puzzles/1", strings.NewReader(`{"name":"test updated","posX":1,"posY":2,"posZ":3,"solution":"test","solved":true}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    body, _ = io.ReadAll(resp.Body)
    assert.Contains(t, string(body), `"name":"test updated"`)

    // Test DeletePuzzle by sending a DELETE request for the puzzle we created.
    req = httptest.NewRequest(http.MethodDelete, "/puzzles/1", nil)
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    // Test GetPuzzles again by sending a GET request for the deleted puzzle. This should return a 404 status.
    req = httptest.NewRequest(http.MethodGet, "/puzzles/1", nil)
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}