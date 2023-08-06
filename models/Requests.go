package models

type PatchRequest struct {
	ID      string `json:"id"`
	Age     int    `json:"age"`
	Surname string `json:"surname"`
}
