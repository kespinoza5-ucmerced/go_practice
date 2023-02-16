package main

import (
	"fmt"

	"github.com/kespinoza5-ucmerced/go_practice/database"
	"github.com/kespinoza5-ucmerced/go_practice/student"
)

func test_student(name, id string) {
	s := student.NewStudent(name, id)

	fmt.Println(*s)

	s.Add_course("CSCI780", 16, 3.70)
	s.Add_course("CSE177", 16, 2.87)
	s.Add_course("CSE195", 16, 4.0)
	s.Add_course("ENGR191", 4, 4.0)
	s.Add_course("WRI105", 16, 3.52)

	fmt.Printf("This student's GPA is: %f\n", s.Calculate_gpa())
}

func test_database(name, id string) {
	db := database.Database{}
	db.Add_student(name, id)

	fmt.Println(db)

	db.Students[0].Add_course("CSCI780", 16, 3.70)
	db.Students[0].Add_course("CSE177", 16, 2.87)
	db.Students[0].Add_course("CSE195", 16, 4.0)
	db.Students[0].Add_course("ENGR191", 4, 4.0)
	db.Students[0].Add_course("WRI105", 16, 3.52)

	// fmt.Println(db)
	// fmt.Println(db.Students[0])
	fmt.Printf("This student's GPA is: %f\n", db.Students[0].Calculate_gpa())

	fmt.Println(db.Find_student_by_id("100741789"))
	fmt.Println(db.Find_student_by_name("Kevin"))

	db.Add_student("Rocky", "100629812")
	db.Add_student("Biggie", "100331265")

	db.Students[1].Add_course("CSE177", 16, 3.40)
	db.Students[1].Add_course("CSE176", 16, 3.69)
	db.Students[2].Add_course("ENGR191", 16, 4.0)
	db.Students[2].Add_course("CSE176", 16, 3.92)

	fmt.Println()
	fmt.Println("Students taking CSE177: ")
	fmt.Println(db.Find_students_by_course("CSE177"))

	fmt.Println()
	fmt.Println("Students taking CSE176: ")
	fmt.Println(db.Find_students_by_course("CSE176"))

	fmt.Println()
	fmt.Printf("Num students in db: %d\n", db.Num_student())
}

func main() {
	name := "Kevin"
	id := "100741789"

	// test_student(name, id)
	test_database(name, id)

}
