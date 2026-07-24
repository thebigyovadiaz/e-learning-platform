package main

// Student represents a learner on the platform
type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewStudent creates and returns a pointer to a new Student
func NewStudent(id, name, email string) *Student {
	return &Student{id, name, email}
}
