package patientstore

import "github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"

type PatientStoreInterface interface {
	Create(p domain.Patient) (*domain.Patient, error)
	GetById(id int) (*domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	UpdateOne(id int, p domain.Patient) (*domain.Patient, error)
	UpdateMany(id int, p domain.Patient) (*domain.Patient, error)
	Delete(id int) error
}