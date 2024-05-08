package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestVerifySolutionHandlerWithIncorrectSolution(t *testing.T) {
    app := fiber.New()
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
    }
}

func TestVerifySolutionHandlerWithInvalidRequest(t *testing.T) {
    app := fiber.New()
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
    }
}

func TestVerifySolutionHandlerWithNoSolutionFound(t *testing.T) {
    app := fiber.New()
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
    }
}
