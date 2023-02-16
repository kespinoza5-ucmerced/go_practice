package student

type course struct {
	name  string
	hours int64
	grade float64
}

type Student struct {
	name, id, major string
	age             int64
	courses         course
}
