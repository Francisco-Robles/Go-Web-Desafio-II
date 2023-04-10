package patient

import (
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
	patientstore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/patient_store"
)

type IPatientRepository interface {
	Create(p domain.Patient) (*domain.Patient, error)
	GetById(id int) (*domain.Patient, error)
	GetAll() ([]domain.Patient, error)
	UpdateOne(id int, p domain.Patient) (*domain.Patient, error)
	UpdateMany(id int, p domain.Patient) (*domain.Patient, error)
	Delete(id int) error
}

type PatientRepository struct {
	Store patientstore.PatientStoreInterface
}

func (pr *PatientRepository) Create(p domain.Patient) (*domain.Patient, error) {

	patient, err := pr.Store.Create(p)
	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (pr *PatientRepository) GetById(id int) (*domain.Patient, error) {

	patient, err := pr.Store.GetById(id)
	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (pr *PatientRepository) GetAll() ([]domain.Patient, error) {

	patients, err := pr.Store.GetAll()
	if err != nil {
		return nil, err
	}

	return patients, nil

}

func (pr *PatientRepository) UpdateOne(id int, p domain.Patient) (*domain.Patient, error) {

	patient, err := pr.Store.UpdateOne(id, p)
	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (pr *PatientRepository) UpdateMany(id int, p domain.Patient) (*domain.Patient, error) {

	patient, err := pr.Store.UpdateMany(id, p)
	if err != nil {
		return nil, err
	}

	return patient, nil

}

func (pr *PatientRepository) Delete(id int) error {

	if err := pr.Store.Delete(id); err != nil {
		return err
	}

	return nil

}