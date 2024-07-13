package member

import "time"

type Response struct {
	ID        uint64    `json:"id,omitempty"`
	Fullname  string    `json:"fullname,omitempty"`
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type RegisterRequest struct {
	ID       string `json:"id" validate:"required,max=100"`
	Username string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
