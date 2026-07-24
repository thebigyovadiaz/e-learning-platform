package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Helper function to read a line
	readLine := func() string {
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

	// Read course information
	courseID := readLine()
	courseTitle := readLine()
	courseInstructor := readLine()

	// Create the course using NewCourse
	course := NewCourse(courseID, courseTitle, courseInstructor)

	// Read number of lessons
	numLessons, _ := strconv.Atoi(readLine())

	// Read each lesson and add to course
	for i := 0; i < numLessons; i++ {
		lessonID := readLine()
		lessonTitle := readLine()
		lessonDuration, _ := strconv.Atoi(readLine())

		// Create lesson and add to course
		lesson := NewLesson(lessonID, lessonTitle, lessonDuration)
		course.AddLesson(lesson)
	}

	// Read student information
	studentID := readLine()
	studentName := readLine()
	studentEmail := readLine()

	// Create the student using NewStudent
	student := NewStudent(studentID, studentName, studentEmail)

	// Create enrollment linking student to course
	enrollment := NewEnrollment(student, course)

	// Read number of operations
	numOps, _ := strconv.Atoi(readLine())

	fmt.Println("")

	// Process each operation
	for i := 0; i < numOps; i++ {
		operation := readLine()

		if operation == "progress" {
			// Call Progress() and print with format "%.1f%%"
			fmt.Printf("%.1f%%\n", enrollment.Progress())
		} else if strings.HasPrefix(operation, "complete ") {
			lessonID := strings.TrimPrefix(operation, "complete ")
			// Call CompleteLesson and print the result
			result := enrollment.CompleteLesson(lessonID)
			fmt.Println(result)
		}
	}
}
