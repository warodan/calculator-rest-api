package models

type UserRequest struct {
	Token        string `json:"token"`
	FirstNumber  int    `json:"first_number"`
	SecondNumber int    `json:"second_number"`
}

type ServerResponse struct {
	Result int `json:"result"`
}
