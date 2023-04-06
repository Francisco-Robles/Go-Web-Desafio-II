package domain

type Dentist struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	License string `json:"license"`
}