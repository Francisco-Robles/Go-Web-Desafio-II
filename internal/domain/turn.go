package domain

type Turn struct {
	Id int `json:"id"`
	DateTime string `json:"datetime"`
	Description string `json:"description"`
	PatientId int `json:"patient_id"`
	DentistId int `json:"dentist_id"`
}

type TurnDTO struct {
	Id int `json:"id"`
	DateTime string `json:"datetime"`
	Description string `json:"description"`
	Patient PatientDTO `json:"Patient"`
	Dentist DentistDTO `json:"Dentist"`
}