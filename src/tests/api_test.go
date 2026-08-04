package tests

import (
	// "main/src/controllers"
	"main/src/routes"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v3"

	// "strings"
	"bytes"
	"encoding/json"
	"main/src/db"
	"os"
	"testing"
)

var testUser = struct{
	Email string `json:"email"`
	Password string `json:"password"`
}{
	Email: "testauto7@gmail.com",
	Password: "passwordauto7",
}

func TestMain(m *testing.M) {

	db.ConnectDB()

	code := m.Run()

	os.Exit(code)
}

var jwtCookie *http.Cookie

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

	cookies := resp.Cookies()

	jwtCookie = cookies[0]

}

func TestLogout(t *testing.T){
	app := fiber.New()

	routes.AuthRoutes(app)

	req := httptest.NewRequest(
		"POST",
		"/auth/logout",
		nil,
	)

	req.AddCookie(jwtCookie)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Error("expected 200")
	}
}