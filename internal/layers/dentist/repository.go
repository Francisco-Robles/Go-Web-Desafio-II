package dentist

import (
	"fmt"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
	dentiststore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/dentist_store"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/web"
)

type IDentistRepository interface {
	Create(d domain.Dentist) (*domain.Dentist, error)
	GetById(id int) (*domain.Dentist, error)
	GetAll() ([]domain.Dentist, error)
	UpdateOne(id int, d domain.Dentist) (*domain.Dentist, error)
	UpdateMany(id int, d domain.Dentist) (*domain.Dentist, error)
	Delete(id int) error
}

type DentistRepository struct {
	Store dentiststore.DentistStoreInterface
}

func (dr *DentistRepository) Create(d domain.Dentist) (*domain.Dentist, error) {

	dentist, err := dr.Store.Create(d)

	if err != nil {
		return nil, web.NewBadRequestApiError(fmt.Sprintf("bad request."))
	}

	return dentist, nil

}

func (dr *DentistRepository) GetById(id int) (*domain.Dentist, error) {

	dentist, err := dr.Store.GetById(id)

	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("dentist id = %d not found", id))
	}

	return dentist, nil

}

func (dr *DentistRepository) GetAll() ([]domain.Dentist, error) {

	dentists, err := dr.Store.GetAll()

	if err != nil {
		return nil, web.NewBadRequestApiError(fmt.Sprintf("bad request."))
	}

	return dentists, nil

}

func (dr *DentistRepository) UpdateOne(id int, d domain.Dentist) (*domain.Dentist, error) {

	dentist, err := dr.Store.UpdateOne(id, d)

	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("dentist id = %d not found", id))
	}

	return dentist, nil

}

func (dr *DentistRepository) UpdateMany(id int, d domain.Dentist) (*domain.Dentist, error) {

	dentist, err := dr.Store.UpdateOne(id, d)

	if err != nil {
		return nil, web.NewNotFoundApiError(fmt.Sprintf("dentist id = %d not found", id))
	}

	return dentist, nil

}

func (dr *DentistRepository) Delete(id int) error {

	if err := dr.Store.Delete(id); err != nil {
		return web.NewNotFoundApiError(fmt.Sprintf("dentist id = %d not found", id))
	}

	return nil

}