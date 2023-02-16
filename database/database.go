package database

import "github.com/kespinoza5-ucmerced/go_practice/student"

type Database struct {
	Students []student.Student
}

// Add_student adds a student to database and returns a pointer to it
func (d *Database) Add_student(name, id string) *student.Student {
	return &student.Student{}
}

// Find_student_by_id returns object with matching student id, otherwise returns nil
func (d *Database) Find_student_by_id(id string) student.Student {
	return student.Student{}
}

// Find_student_by_name returns object with matching student name, otherwise returns nil
// Assumes student's name is unique
func (d *Database) Find_student_by_name(id string) student.Student {
	return student.Student{}
}

// Find_student_by_id returns list of students taking the course, otherwise returns nil
func (d *Database) Find_students_by_course(course_name string) []student.Student {
	return []student.Student{}
}

// Num_student
func (d *Database) Num_student() int {
	return 0
}
