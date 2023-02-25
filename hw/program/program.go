package program

import (
	"fmt"
	"strconv"

	"github.com/kespinoza5-ucmerced/go_practice/hw/database"
	"github.com/kespinoza5-ucmerced/go_practice/hw/student"
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

type Program struct {
	userIO               string
	database             database.Database
	last_student         *student.Student
	transaction_executed int64
}

// NewProgram constructs new program with user provided userIO and database objects
func NewProgram(_userio string, _database database.Database) *Program {
	return &Program{
		database: _database,
	}
}

func (p *Program) add_student_to_db() {
	var s_name, s_id, s_major string
	var s_age int64

	fmt.Println("Enter student's name, id, major, and age:")

	fmt.Scanln(&s_name, &s_id, &s_major, &s_age)

	p.last_student = p.database.Add_student(s_name, s_id)
	p.last_student.Major = s_major
	p.last_student.Age = s_age

	p.transaction_executed += 1

	fmt.Println("Student " + p.last_student.Id + " added\n")
}

func (p *Program) Run() {
	program_menu :=
		"1. Add a student\n" +
			"2. Find a student by id\n" +
			"3. Find a student by name\n" +
			"4. List students by course\n" +
			"5. Add a course to a student\n" +
			"6. Print the number of transactions executed\n" +
			"7. Exit\n"

	// init menu object

	for {
		fmt.Println(program_menu)

		var input string
		fmt.Scanln(&input)
		input_num, _ := strconv.Atoi(input)

		// set menu object

		// run user_input from menu object

		if input_num == EXIT {
			break
		} else if input_num == ADD_STUDENT {
			var s_name, s_id, s_major string
			var s_age int64

			fmt.Println("Enter student's name, id, major, and age:")
			fmt.Scanln(&s_name, &s_id, &s_major, &s_age)

			p.last_student = p.database.Add_student(s_name, s_id)
			p.last_student.Major = s_major
			p.last_student.Age = s_age

			fmt.Printf("Student %s added\n", p.last_student.Id)

			p.transaction_executed += 1
		} else if input_num == FIND_STUDENT_BY_ID {
			var s_id string

			fmt.Println("Enter student's id:")
			fmt.Scanln(&s_id)

			student := p.database.Find_student_by_id(s_id)

			if student == nil {
				fmt.Println("Student not found")
				student = p.last_student
			} else {
				output :=
					"Student\n" +
						"\tName: " + p.last_student.Name + "\n" +
						"\tID: " + p.last_student.Id + "\n" +
						"\tMajor: " + p.last_student.Major + "\n" +
						"\tAge: " + strconv.Itoa(int(p.last_student.Age))

				fmt.Println()
				fmt.Println(output)
				fmt.Println()
			}

			p.transaction_executed += 1
		} else if input_num == FIND_STUDENT_BY_NAME {
			var s_name string

			fmt.Println("Enter student's name:")
			fmt.Scanln(&s_name)

			student := p.database.Find_student_by_name(s_name)

			if student == nil {
				fmt.Println("Student not found")
				student = p.last_student
			} else {
				output :=
					"Student\n" +
						"\tName: " + p.last_student.Name + "\n" +
						"\tID: " + p.last_student.Id + "\n" +
						"\tMajor: " + p.last_student.Major + "\n" +
						"\tAge: " + strconv.Itoa(int(p.last_student.Age))

				fmt.Println()
				fmt.Println(output)
				fmt.Println()
			}

			p.transaction_executed += 1
		} else if input_num == LIST_STUDENTS_BY_COURSE {
			var c_name string

			fmt.Println("Enter course name:")
			fmt.Scanln(&c_name)

			students := p.database.Find_students_by_course(c_name)

			if students == nil {
				fmt.Println("No student found")
			} else {
				output := "["
				for i, student := range students {
					output += student.Name
					if i != len(students)-1 {
						output += ", "
					}
				}
				output += "]"

				println(output)
			}

			p.transaction_executed += 1
		} else if input_num == ADD_COURSE_TO_STUDENT {
			if p.last_student == nil {
				fmt.Println("Search or add student before adding a course")
			} else {
				var c_name string
				var c_hours int64
				var c_grade float64

				fmt.Println("Enter course name, credit hours, and grade:")
				fmt.Scanln(&c_name, &c_hours, &c_grade)
				p.last_student.Add_course(c_name, c_hours, c_grade)
				fmt.Printf("Course %s added to student %s", c_name, p.last_student.Name)
				p.transaction_executed += 1
			}
		} else if input_num == PRINT_NUM_TX {
			fmt.Printf("Number of transactions: %d\n", p.transaction_executed)
			p.transaction_executed += 1
		}
	}
}
