package student

type course struct {
	name  string
	hours int64
	grade float64
}

type Student struct {
	Name, Id, Major string
	Age             int64
	courses         []course
}

// NewStudent initializes Student with given name and id
func NewStudent(_name, _id string) *Student {
	return &Student{
		Name: _name,
		Id:   _id,
	}
}

// Add_course takes args and pushes into that student's course list
func (s *Student) Add_course(course_name string, credit_hours int64, grade float64) {
	s.courses = append(s.courses, course{course_name, credit_hours, grade})
}

// Calculate_gpa returns gpa for student across all courses in their course list
// GPA = sum(grade * credit_hours) / sum(credit_hours)
func (s *Student) Calculate_gpa() float64 {
	if len(s.courses) == 0 {
		return 0
	}

	var credit_sum, grade_credit_sum float64

	for _, course := range s.courses {
		credit_sum += float64(course.hours)
		grade_credit_sum += course.grade * float64(course.hours)
	}

	return grade_credit_sum / credit_sum
}
