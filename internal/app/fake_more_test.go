package app_test

import (
	"github.com/hjyoon/socketing-server/internal/api"
	"github.com/hjyoon/socketing-server/internal/store"
)

func (f *fakeStore) CreateEvent(string, store.EventInput) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) UpdateEvent(string, store.EventInput) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) DeleteEvent(string) api.Error { return api.NoError }
func (f *fakeStore) CreateAreas(string, store.AreaBatch) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) ListSeats(string) ([]map[string]any, api.Error) {
	return []map[string]any{ok()}, api.NoError
}
func (f *fakeStore) GetSeat(string, string) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) SeatStatus(string, string, string) (any, api.Error) {
	return []map[string]any{ok()}, api.NoError
}
func (f *fakeStore) ListOrders(string, string) ([]map[string]any, api.Error) {
	return []map[string]any{ok()}, api.NoError
}
func (f *fakeStore) GetOrder(string, string) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) CancelOrder(string, string) api.Error { return api.NoError }
func (f *fakeStore) CreatePayment(string, store.PaymentInput) (map[string]any, api.Error) {
	return ok(), api.NoError
}
func (f *fakeStore) ManagerEvent(string, string, string) (map[string]any, api.Error) {
	return ok(), api.NoError
}
