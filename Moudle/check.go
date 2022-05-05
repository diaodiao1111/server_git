package Moudle

type check struct {
	Name   string
	Age    int
	Genger string
	id     int
	Score  float64
}

func NewCheck(n string, age int, ID int) *check {
	return &check{
		Name: n,
		Age:  age,
		id:   ID,
	}
}

func (c *check) GetCheckID() int {
	return c.id
}
