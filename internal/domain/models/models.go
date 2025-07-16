package models

type UserRequest struct {
	Token        string `json:"token" validate:"required"`
	FirstNumber  *int   `json:"first_number" validate:"required"`
	SecondNumber *int   `json:"second_number" validate:"required"`
}

type ServerResponse struct {
	Result int `json:"result"`
}
