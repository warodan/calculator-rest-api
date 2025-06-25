package models

type SumRequest struct {
	Token        string `json:"token"`
	FirstNumber  int    `json:"first_number"`
	SecondNumber int    `json:"second_number"`
}

type SumResponse struct {
	Result int `json:"result"`
}
