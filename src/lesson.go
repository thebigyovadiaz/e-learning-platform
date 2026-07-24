package main

// Lesson represent an individual learning unit
type Lesson struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
}

// NewLesson creates and returns a pointer to a New Lesson
func NewLesson(id, title string, duration int) *Lesson {
	return &Lesson{id, title, duration}
}
