package main

// Progressable interface for types that can report completion progress
type Progressable interface {
	Progress() float64
}

// Enrollment links a student to a course and tracks completed lessons
type Enrollment struct {
	Student          *Student
	Course           *Course
	CompletedLessons map[string]bool
}

func NewEnrollment(student *Student, course *Course) *Enrollment {
	return &Enrollment{student, course, make(map[string]bool)}
}

// CompleteLesson marks a lesson as complete
// Returns "lesson not found" if lesson doesn't exist in course
// Returns "already completed" if lesson was already completed
// Returns "completed" on success
func (e *Enrollment) CompleteLesson(lessonID string) string {
	// Check if lesson exist in the course
	found := false
	for _, lesson := range e.Course.Lessons {
		if lesson.ID == lessonID {
			found = true
			break
		}
	}

	if !found {
		return "lesson not found"
	}

	// Check if lesson is already completed
	if e.CompletedLessons[lessonID] {
		return "already completed"
	}

	// Mark lesson as completed
	e.CompletedLessons[lessonID] = true
	return "lesson completed"
}

// Progress returns the percentage of lessons completed (0.0 to 100.0)
func (e *Enrollment) Progress() float64 {
	totalLessons := len(e.Course.Lessons)
	if totalLessons == 0 {
		return 0.0
	}

	completedCount := len(e.CompletedLessons)
	return float64(completedCount) / float64(totalLessons) * 100.0
}
