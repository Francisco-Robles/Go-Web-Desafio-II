package domain

type Patient struct {
	Id            int `json:"id"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Address       string `json:"address"`
	Dni           string `json:"dni"`
	DischargeDate string `json:"discharge_date"`
}

type PatientDTO struct {
	Name	string `json:"name"`
	Surname string `json:"surname"`
	Address string `json:"address"`
	Dni string  `json:"dni"`
}