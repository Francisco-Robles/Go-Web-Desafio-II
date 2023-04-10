package dentiststore

import "github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"

type DentistStoreInterface interface {
	Create(d domain.Dentist) (*domain.Dentist, error)
	GetById(id int) (*domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	UpdateOne(id int, d domain.Dentist) (*domain.Dentist, error)
	UpdateMany(id int, d domain.Dentist) (*domain.Dentist, error)
	Delete(id int) error
	GetIdByLicense(license string) (int, error)
}