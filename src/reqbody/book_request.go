package reqbody

type SaveBookRequest struct {
	Id int64 `json:"id"` // to update, include id.
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Description string `json:"description" binding:"max=50"`
	TotalStock int32 `json:"total_stock" binding:"required"`
	Content string `json:"content"`
}

