package api

type Error struct {
	Code    int
	Message string
	Status  int
}

var (
	NoError           = Error{}
	ErrUserExists     = Error{1, "User already exists", 409}
	ErrCredentials    = Error{2, "Invalid credentials", 401}
	ErrNicknameExists = Error{4, "Nickname already exists", 409}
	ErrValidation     = Error{5, "Validation failed", 400}
	ErrInternal       = Error{6, "Internal server error", 500}
	ErrUserNotFound   = Error{7, "User not found", 404}
	ErrUnauthorized   = Error{8, "Unauthorized", 401}
	ErrEventNotFound  = Error{9, "Event not found", 404}
	ErrDuplicateSeat  = Error{10, "A seat with the same area, row, number, and event already exists", 409}
	ErrSeatNotFound   = Error{11, "Seat not found for the specified event", 404}
	ErrEventDate      = Error{12, "Event date not found for the specified event", 404}
	ErrReservation    = Error{13, "This seat is already reserved for the selected event date", 409}
	ErrOrderNotFound  = Error{15, "Order not found", 404}
	ErrExistingOrder  = Error{17, "The order is already pending or completed for the reservations", 409}
	ErrBalance        = Error{20, "This payment is not acceptable because of insufficient balance", 409}
	ErrCanceledOrder  = Error{22, "This order has already been canceled. Further cancellation is not allowed.", 409}
)
