package models

type SumRequest struct {
	FirstNumber  int `json:"first_number"`
	SecondNumber int `json:"second_number"`
}

type SumResponse struct {
	Result int `json:"result"`
}
