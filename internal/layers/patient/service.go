package patient

import (

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
)

type IPatientService interface {
	Create(p domain.Patient) (*domain.Patient, error)
	GetById(id int) (*domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	UpdateOne(id int, p domain.Patient) (*domain.Patient, error)
	UpdateMany(id int, p domain.Patient) (*domain.Patient, error)
	Delete(id int) error
}

type PatientService struct {
	Repository IPatientRepository
}

func (ps *PatientService) Create(p domain.Patient) (*domain.Patient, error) {

	patient, err := ps.Repository.Create(p)

	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (ps *PatientService) GetById(id int) (*domain.Patient, error) {

	patient, err := ps.Repository.GetById(id)

	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (ps *PatientService) GetAll() ([]domain.Patient, error) {

	patients, err := ps.Repository.GetAll()

	if err != nil {
		return nil, err
	}

	return patients, nil

}

func (ps *PatientService) UpdateOne(id int, p domain.Patient) (*domain.Patient, error) {

	patient, err := ps.Repository.UpdateOne(id, p)

	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (ps *PatientService) UpdateMany(id int, p domain.Patient) (*domain.Patient, error) {

	patient, err := ps.Repository.UpdateMany(id, p)

	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (ps *PatientService) Delete(id int) error {

	if err := ps.Repository.Delete(id); err != nil {
		return err
	}

	return nil

}