package store

import "github.com/hjyoon/socketing-server/internal/api"

type Store interface {
	Health() error
	EnsureSchema() error
	CreateUser(Register) (map[string]any, api.Error)
	Login(Login, string) (map[string]any, api.Error)
	GetUser(id string) (map[string]any, api.Error)
	GetUserByEmail(email string) (map[string]any, api.Error)
	GetPoints(id string) (map[string]any, api.Error)
	UpdateNickname(id, nickname string) (map[string]any, api.Error)
	UpdatePassword(id, password string) api.Error
	DeleteUser(id string) api.Error
	ListEvents(userID string) ([]map[string]any, api.Error)
	GetEvent(id string) (map[string]any, api.Error)
	CreateEvent(string, EventInput) (map[string]any, api.Error)
	UpdateEvent(string, EventInput) (map[string]any, api.Error)
	DeleteEvent(id string) api.Error
	CreateAreas(string, AreaBatch) (map[string]any, api.Error)
	ListSeats(eventID string) ([]map[string]any, api.Error)
	GetSeat(eventID, seatID string) (map[string]any, api.Error)
	SeatStatus(eventID, dateID, seatID string) (any, api.Error)
	ListOrders(userID, eventID string) ([]map[string]any, api.Error)
	GetOrder(orderID, userID string) (map[string]any, api.Error)
	CancelOrder(orderID, userID string) api.Error
	CreatePayment(userID string, p PaymentInput) (map[string]any, api.Error)
	ManagerEvent(userID, eventID, dateID string) (map[string]any, api.Error)
}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EventDateInput struct {
	Date string `json:"date"`
}

type EventInput struct {
	Title              string           `json:"title"`
	Thumbnail          string           `json:"thumbnail"`
	Place              string           `json:"place"`
	Cast               string           `json:"cast"`
	AgeLimit           *int             `json:"ageLimit"`
	SVG                string           `json:"svg"`
	TicketingStartTime string           `json:"ticketingStartTime"`
	EventDates         []EventDateInput `json:"eventDates"`
}

type SeatInput struct {
	Cx     int `json:"cx"`
	Cy     int `json:"cy"`
	Row    int `json:"row"`
	Number int `json:"number"`
}

type AreaInput struct {
	Label string      `json:"label"`
	Price int         `json:"price"`
	SVG   string      `json:"svg"`
	Seats []SeatInput `json:"seats"`
}

type AreaBatch struct {
	Areas []AreaInput `json:"areas"`
}

type PaymentInput struct {
	PaymentMethod string   `json:"paymentMethod"`
	EventDateID   string   `json:"eventDateId"`
	SeatIDs       []string `json:"seatIds"`
}
