package app_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hjyoon/socketing-server/internal/app"
	"github.com/hjyoon/socketing-server/internal/auth"
)

func TestRouterPublicRoutes(t *testing.T) {
	st := newFake()
	r := app.NewRouter(app.Config{JWTSecret: "s"}, st)
	cases := []struct {
		method string
		path   string
		body   string
		code   int
	}{
		{"GET", "/api/health", "", 200},
		{"POST", "/api/auth/register", `{"email":"a","password":"p"}`, 201},
		{"POST", "/api/auth/login", `{"email":"a","password":"p"}`, 200},
		{"GET", "/api/users/email/a", "", 200},
		{"GET", "/api/users/u/points", "", 200},
		{"GET", "/api/users/u", "", 200},
		{"GET", "/api/events/", "", 200},
		{"GET", "/api/events/e", "", 200},
		{"GET", "/api/events/e/seats", "", 200},
		{"GET", "/api/events/e/seats/s", "", 200},
	}
	for _, tc := range cases {
		res := req(r, tc.method, tc.path, tc.body, "")
		if res.Code != tc.code {
			t.Fatalf("%s %s got %d", tc.method, tc.path, res.Code)
		}
	}
}

func TestRouterProtectedRoutes(t *testing.T) {
	st := newFake()
	r := app.NewRouter(app.Config{JWTSecret: "s"}, st)
	token := auth.Sign("u", "s")
	cases := []struct {
		method string
		path   string
		body   string
		code   int
	}{
		{"PATCH", "/api/users/u/nickname", `{"nickname":"n"}`, 200},
		{"PATCH", "/api/users/u/password", `{"password":"p"}`, 200},
		{"DELETE", "/api/users/u", "", 204},
		{"POST", "/api/events/", `{"title":"t"}`, 201},
		{"PUT", "/api/events/e", `{"title":"t"}`, 200},
		{"DELETE", "/api/events/e", "", 204},
		{"POST", "/api/events/e/seats/batch", `{"areas":[]}`, 201},
		{"POST", "/api/events/e/seats/batch", `{`, 400},
		{"GET", "/api/events/e/seats-status", "", 200},
		{"GET", "/api/events/e/seats-status/s", "", 200},
		{"GET", "/api/orders/", "", 201},
		{"GET", "/api/orders/o", "", 201},
		{"POST", "/api/orders/o/cancel", "", 201},
		{"POST", "/api/payments/", `{"seatIds":[]}`, 201},
		{"POST", "/api/payments/", `{`, 400},
		{"PATCH", "/api/payments/", `{}`, 201},
		{"GET", "/api/managers/events/", "", 200},
		{"GET", "/api/managers/events/e/reservation-status", "", 200},
	}
	for _, tc := range cases {
		res := req(r, tc.method, tc.path, tc.body, token)
		if res.Code != tc.code {
			t.Fatalf("%s %s got %d", tc.method, tc.path, res.Code)
		}
	}
}

func TestRouterCORSAllowsAuthorization(t *testing.T) {
	st := newFake()
	r := app.NewRouter(app.Config{
		JWTSecret:   "s",
		CORSOrigins: []string{"http://localhost:5173"},
	}, st)
	req := httptest.NewRequest("OPTIONS", "/api/events/", nil)
	req.Header.Set("Origin", "http://localhost:5173")
	req.Header.Set("Access-Control-Request-Method", "GET")
	req.Header.Set("Access-Control-Request-Headers", "authorization,content-type")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 204 {
		t.Fatalf("preflight got %d", w.Code)
	}
	headers := w.Header().Get("Access-Control-Allow-Headers")
	if !strings.Contains(headers, "Authorization") {
		t.Fatalf("missing Authorization in %q", headers)
	}
}

func req(h http.Handler, method, path, body, token string) *httptest.ResponseRecorder {
	r := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	r.Header.Set("Content-Type", "application/json")
	if token != "" {
		r.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	return w
}
