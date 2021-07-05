package model

type Film struct {
	Title  string	`json:"title"`
	Year   int		`json:"year"`
	Rating float64	`json:"rating"`
	Link   string	`json:"link"`
}