package programmenu

import (
	"fmt"

	"github.com/kespinoza5-ucmerced/go_practice/hw/database"
)

const (
	ADD_STUDENT             = 1
	FIND_STUDENT_BY_ID      = 2
	FIND_STUDENT_BY_NAME    = 3
	LIST_STUDENTS_BY_COURSE = 4
	ADD_COURSE_TO_STUDENT   = 5
	PRINT_NUM_TX            = 6
	EXIT                    = 7
)

type menu_action interface {
	action()
	get_menu_text()
}

type action_add_student struct {
	menu_text string
}

func New_action_add_student() *action_add_student {
	return &action_add_student{
		menu_text: "Add a student",
	}
}

func (a *action_add_student) get_menu_text() string {
	return a.menu_text
}

func (a *action_add_student) process(d *database.Database) int {
	var s_name, s_id, s_major string
	var s_age int64

	fmt.Println("Enter student's name, id, major, and age:")

	fmt.Scanln(&s_name, &s_id, &s_major, &s_age)

	student := d.Add_student(s_name, s_id)
	student.Major = s_major
	student.Age = s_age

	return 0
}

type action_find_student_by_id struct {
}

type action_find_student_by_name struct {
}

type action_list_students_by_course struct {
}

type action_add_course_to_student struct {
}

type action_print_num_tx struct {
}

type Menu struct {
	actions []menu_action
}
