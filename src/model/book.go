package model
import "fmt"

type Book struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	Author string `json:"author" binding:"required"`
	Description string `json:"description" binding:"max=50"`
	TotalStock int32 `json:"total_stock"`
	OnLoan int32 `json:"on_loan"`
}

type BookContent struct {
	Id int64 `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (b Book) String() string {
	return fmt.Sprintf("[%d, %s, %s,\n %s \n Total: %d, On loan %d, Available: %d]", b.Id, b.Title, b.Author, b.Description, b.TotalStock, b.OnLoan, b.TotalStock - b.OnLoan)
}