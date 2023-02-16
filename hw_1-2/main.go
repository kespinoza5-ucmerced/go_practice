package main

import (
	"fmt"

	"github.com/kespinoza5-ucmerced/go_practice/database"
	"github.com/kespinoza5-ucmerced/go_practice/student"
)

func main() {
	name := "Kevin"
	id := "100741789"

	s := student.NewStudent(name, id)

	fmt.Println(*s)

	s.Add_course("CSCI780", 16, 3.70)
	s.Add_course("CSE177", 16, 2.87)
	s.Add_course("CSE195", 16, 4.0)
	s.Add_course("ENGR191", 4, 4.0)
	s.Add_course("WRI105", 16, 3.52)

	fmt.Printf("This student's GPA is: %f\n", s.Calculate_gpa())

	db := database.Database{}
	db.Add_student(name, id)

	fmt.Println(db)

	// db.Students[0].Add_course("CSCI780", 16, 3.70)
	// db.Students[0].Add_course("CSE177", 16, 2.87)
	// db.Students[0].Add_course("CSE195", 16, 4.0)
	// db.Students[0].Add_course("ENGR191", 4, 4.0)
	// db.Students[0].Add_course("WRI105", 16, 3.52)

	// fmt.Printf("This student's GPA is: %f\n", db.Students[0].Calculate_gpa())
}
