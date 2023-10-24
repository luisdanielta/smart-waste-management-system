package models

type Data_t struct {
	X int
	Y int
	Z int
}

type Container struct {
	Name     string
	Location Location_t
	Usecase  string // School, park, etc.
	Size     string // S, M, L - small, medium, large
	Data     Data_t
}
