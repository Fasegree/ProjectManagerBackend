package models

type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Field       string `json:"field"`
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Experience  string `json:"experience"`
}
