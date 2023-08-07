package models

type Person struct {
	ID      string `json:"id"`
	Name    string `json:"name" `
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}
