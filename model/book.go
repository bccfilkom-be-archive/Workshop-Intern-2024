package model

type CreateBook struct {
	Title       string `json:"title" binding:"required"`
	Writter     string `json:"writter" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       uint   `json:"stock" binding:"required"`
}

type UpdateBook struct {
	Title       string `json:"title"`
	Writter     string `json:"writter"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
}
