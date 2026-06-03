package app_test

import (
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

type fakeStore struct{}

func newFake() *fakeStore { return &fakeStore{} }

func ok() map[string]any { return map[string]any{"id": "x"} }

func (f *fakeStore) Health() error       { return nil }
func (f *fakeStore) EnsureSchema() error { return nil }
func (f *fakeStore) CreateUser(store.Register) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) Login(store.Login, string) (map[string]any, api.Error) {
	return map[string]any{"accessToken": "t"}, api.NoError
}
func (f *fakeStore) GetUser(string) (map[string]any, api.Error)        { return ok(), api.NoError }
func (f *fakeStore) GetUserByEmail(string) (map[string]any, api.Error) { return ok(), api.NoError }
func (f *fakeStore) GetPoints(string) (map[string]any, api.Error) {
	return map[string]any{"point": 1}, api.NoError
}
func (f *fakeStore) UpdateNickname(string, string) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) UpdatePassword(string, string) api.Error { return api.NoError }
func (f *fakeStore) DeleteUser(string) api.Error             { return api.NoError }
func (f *fakeStore) ListEvents(string) ([]map[string]any, api.Error) {
	return []map[string]any{ok()}, api.NoError
}
func (f *fakeStore) GetEvent(string) (map[string]any, api.Error) { return ok(), api.NoError }
