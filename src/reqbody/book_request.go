package reqbody

type SaveBookRequest struct {
	Id int64 `json:"id"` // to update, include id.
	Title string `json:"title" binding:"required,min=4"`
	Author string `json:"author" binding:"required,min=4"`
	Description string `json:"description" binding:"max=50"`
	TotalStock int32 `json:"total_stock" binding:"required"`
	Content string `json:"content"`
}

type FindBookRequest struct {
	BookId int64 `form:"id"`
	Title string `form:"title"`
	Author string `form:"author"`
}