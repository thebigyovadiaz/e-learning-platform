package main

// Course represents a course with multiple lessons
type Course struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Instructor string    `json:"instructor"`
	Lessons    []*Lesson `json:"lessons"`
}

// NewCourse creates and returns a pointer to a new Course
func NewCourse(id, title, instructor string) *Course {
	return &Course{id, title, instructor, make([]*Lesson, 0)}
}

// AddLesson adds a lesson to the course
func (c *Course) AddLesson(l *Lesson) {
	c.Lessons = append(c.Lessons, l)
}

// TotalDuration returns the sum of all lesson durations
func (c *Course) TotalDuration() int {
	total := 0
	for _, l := range c.Lessons {
		total += l.Duration
	}

	return total
}
