package model

type Book struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	TotalStock int32 `json:"total_stock"`
	OnLoan int32 `json:"on_loan"`
}
