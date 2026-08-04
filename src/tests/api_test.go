package tests

import (
	// "main/src/controllers"
	"main/src/routes"
	"net/http/httptest"
 	"github.com/gofiber/fiber/v3"
	// "strings"
	"encoding/json"
	"bytes"
	"testing"
	"main/src/db"
	"os"
)

var testUser = struct{
	Email string `json:"email"`
	Password string `json:"password"`
}{
	Email: "testauto5@gmail.com",
	Password: "passwordauto5",
}

func TestMain(m *testing.M) {

	db.ConnectDB()

	code := m.Run()

	os.Exit(code)
}

func TestRegister(t *testing.T){
	app := fiber.New()

	routes.AuthRoutes(app)

	// app.Post("/register", controllers.RegisterUser)
	body, _ := json.Marshal(testUser)

	req := httptest.NewRequest(
		"POST",
		"/auth/register",
		bytes.NewReader(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Error("expected 200")
	}

}

func TestLogin(t *testing.T){
	app := fiber.New()

	routes.AuthRoutes(app)

	body, _ := json.Marshal(testUser)

	req := httptest.NewRequest(
		"POST",
		"/auth/login",
		bytes.NewReader(body),
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Error("expected 200")
	}
}