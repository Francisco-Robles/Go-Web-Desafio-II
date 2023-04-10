package domain

type Turn struct {
	Id int `json:"id"`
	Description string `json:"description"`
	PatientId int `json:"patient_id"`
	DentistId int `json:"dentist_id"`
}