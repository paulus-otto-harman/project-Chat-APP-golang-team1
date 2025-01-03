package model

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type Pagination struct {
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	TotalItems int       `json:"total_items"`
	TotalPages int       `json:"total_pages"`
	RoomId     int       `json:"room_id"`
	RoomName   string    `json:"room_name"`
	Messages   []Message `json:"data"`
}
