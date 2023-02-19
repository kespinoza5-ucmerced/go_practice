package database

import "github.com/kespinoza5-ucmerced/go_practice/hw/student"

type Database struct {
	Students []student.Student
}

// Add_student adds a student to database and returns a pointer to it
func (d *Database) Add_student(_name, _id string) *student.Student {
	d.Students = append(d.Students, *student.NewStudent(_name, _id))
	new_student_index := len(d.Students) - 1

	return &d.Students[new_student_index]
}

// Find_student_by_id returns object with matching student id, otherwise returns nil
func (d *Database) Find_student_by_id(_id string) *student.Student {
	for _, student := range d.Students {
		if student.Id == _id {
			return &student
		}
	}

	return nil
}

// Find_student_by_name returns object with matching student name, otherwise returns nil
//
// Assumes student's name is unique
func (d *Database) Find_student_by_name(_name string) *student.Student {
	for _, student := range d.Students {
		if student.Name == _name {
			return &student
		}
	}

	return nil
}

// Find_student_by_course returns list of students taking the course, otherwise returns nil
func (d *Database) Find_students_by_course(_course_name string) []student.Student {
	var students_taking_course []student.Student

	for _, student := range d.Students {
		for _, course := range student.Courses {
			if course.Name == _course_name {
				students_taking_course = append(students_taking_course, student)
			}
		}
	}

	if len(students_taking_course) == 0 {
		return nil
	}

	return students_taking_course
}

// Num_student returns number of students in database
func (d *Database) Num_student() int {
	return len(d.Students)
}
