package entity

type Users struct {
	ID      string
	Name    string
	Address string
}

type Student struct {
	ID      string
	Name    string
	Address string
	Grade   int
}

var DataStudent = []Student{
	{"01", "Ali", "", 5},
	{"02", "Adi", "", 6},
}
