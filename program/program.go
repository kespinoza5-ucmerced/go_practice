package program

import (
	"fmt"
	"strconv"

	"github.com/kespinoza5-ucmerced/go_practice/database"
	"github.com/kespinoza5-ucmerced/go_practice/student"
)

type Program struct {
	userIO               string
	database             database.Database
	last_student         student.Student
	transaction_executed int64
}

// NewProgram constructs new program with user provided userIO and database objects
func NewProgram(_userio string, _database database.Database) *Program {
	return &Program{
		userIO:   _userio,
		database: _database,
	}
}

func (p *Program) add_student_to_db() {
	var s_name, s_id, s_major string
	var s_age int64

	fmt.Println("Enter student's name, id, major, and age:")

	fmt.Scanln(&s_name, &s_id, &s_major, &s_age)

	p.database.Add_student(s_name, s_id)
	index := len(p.database.Students)
	p.database.Students[index].Major = s_major
	p.database.Students[index].Age = s_age
	p.last_student = p.database.Students[index]
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

	var input string

	fmt.Println(program_menu)
	fmt.Scanln(&input)
	input_num, _ := strconv.Atoi(input)

	for input_num != 7 {
		if input_num == 1 {
			p.add_student_to_db()
		}

		fmt.Println(program_menu)
		fmt.Scanln(&input)
		input_num, _ = strconv.Atoi(input)
	}
}
