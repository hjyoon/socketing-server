package app_test

import (
	"errors"
	"testing"

	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/app"
)

type failingStore struct{ fakeStore }

func (f failingStore) Health() error { return errors.New("down") }
func (f failingStore) GetUser(string) (map[string]any, api.Error) {
	return nil, api.ErrUserNotFound
}

func TestRouterErrorPaths(t *testing.T) {
	r := app.NewRouter(app.Config{JWTSecret: "s"}, &failingStore{})
	cases := []struct {
		method string
		path   string
		body   string
		code   int
	}{
		{"GET", "/api/health", "", 500},
		{"GET", "/api/users/u", "", 404},
		{"PATCH", "/api/users/u/nickname", `{"nickname":"n"}`, 401},
		{"PATCH", "/api/users/u/nickname", `{"nickname":"n"}`, 401},
		{"POST", "/api/auth/register", `{`, 400},
		{"POST", "/api/auth/login", `{`, 400},
	}
	for _, tc := range cases {
		res := req(r, tc.method, tc.path, tc.body, "")
		if res.Code != tc.code {
			t.Fatalf("%s %s got %d", tc.method, tc.path, res.Code)
		}
	}
	bad := req(r, "POST", "/api/payments/", `{`, "bad")
	if bad.Code != 401 {
		t.Fatalf("bad token got %d", bad.Code)
	}
}
