package structure

type Item struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type UpdateItemInput struct {
	Title       *string `json:"title" binding:"required"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}
