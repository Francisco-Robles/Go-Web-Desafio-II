package dentist

import "github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"

type IDentistService interface {
	Create(d domain.Dentist) (*domain.Dentist, error)
	GetById(id int) (*domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	UpdateOne(id int, d domain.Dentist) (*domain.Dentist, error)
	UpdateMany(id int, d domain.Dentist) (*domain.Dentist, error)
	Delete(id int) error
}

type DentistService struct {
	Repository IDentistRepository
}

func NewDentistService(r IDentistRepository) IDentistService {
	return &DentistService{Repository: r}
}

func (ds *DentistService) Create(d domain.Dentist) (*domain.Dentist, error) {

	dentist, err := ds.Repository.Create(d)
	if err != nil {
		return nil, err
	}

	return dentist, nil

}

func (ds *DentistService) GetById(id int) (*domain.Dentist, error) {

	dentist, err := ds.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return dentist, nil

}

func (ds *DentistService) GetAll() ([]domain.Dentist, error) {

	dentists, err := ds.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return dentists, nil

}

func (ds *DentistService) UpdateOne(id int, d domain.Dentist) (*domain.Dentist, error) {

	dentist, err := ds.Repository.UpdateOne(id, d)
	if err != nil {
		return nil, err
	}

	return dentist, nil

}

func (ds *DentistService) UpdateMany(id int, d domain.Dentist) (*domain.Dentist, error) {

	dentist, err := ds.Repository.UpdateMany(id, d)
	if err != nil {
		return nil, err
	}

	return dentist, nil

}

func (ds *DentistService) Delete(id int) error {

	if err := ds.Repository.Delete(id); err != nil {
		return err
	}

	return nil

}