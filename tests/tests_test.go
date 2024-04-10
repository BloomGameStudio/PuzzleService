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

func TestPuzzleHandler(t *testing.T) {

    app := fiber.New()
    puzzleHandlers := handlers.NewPuzzleHandler()
    app.Get("/puzzles", puzzleHandlers.GetPuzzles)
    app.Post("/puzzles", puzzleHandlers.CreatePuzzle)
    app.Put("/puzzles", puzzleHandlers.UpdatePuzzle)
    app.Delete("/puzzles", puzzleHandlers.DeletePuzzle)

    // Test CreatePuzzle
    req := httptest.NewRequest(http.MethodPost, "/puzzles", strings.NewReader(`{"id":"1","name":"test","posX":1,"posY":2,"posZ":3,"solution":"test","solved":false}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err := app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    // Test GetPuzzles
    req = httptest.NewRequest(http.MethodGet, "/puzzles", nil)
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    body, _ := io.ReadAll(resp.Body)
    assert.Contains(t, string(body), `"id":"1"`)

    // Test UpdatePuzzle
    req = httptest.NewRequest(http.MethodPut, "/puzzles?id=1", strings.NewReader(`{"name":"test updated","posX":1,"posY":2,"posZ":3,"solution":"test","solved":true}`))
    req.Header.Set("Content-Type", "application/json")
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    body, _ = io.ReadAll(resp.Body)
    assert.Contains(t, string(body), `"name":"test updated"`)

    // Test DeletePuzzle
    req = httptest.NewRequest(http.MethodDelete, "/puzzles?id=1", nil)
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)

    // Test GetPuzzles again to ensure the puzzle was deleted
    req = httptest.NewRequest(http.MethodGet, "/puzzles", nil)
    resp, err = app.Test(req)
    assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, resp.StatusCode)
    body, _ = io.ReadAll(resp.Body)
    assert.NotContains(t, string(body), `"id":"1"`)
}