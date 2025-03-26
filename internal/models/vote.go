package models

type Vote struct {
	ID       string
	Question string
	Options  map[int]string
	Author   string
}
