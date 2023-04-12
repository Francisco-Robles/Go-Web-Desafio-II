package turn

import (

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
	turnstore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/turn_store"
)

type ITurnRepository interface {
	Create(t domain.Turn) (*domain.Turn, error)
	GetById(id int) (*domain.TurnDTO, error)
	FindById(id int)  (*domain.Turn, error)
	GetAll() ([]domain.Turn, error)
	UpdateOne(id int, t domain.Turn) (*domain.Turn, error)
	UpdateMany(id int, t domain.Turn) (*domain.Turn, error)
	Delete(id int) error
	CreateByPatientDniAndDentistLicense(t domain.Turn, dni string, license string) (*domain.Turn, error)
	GetByPatientDni(dni string) (*domain.TurnDTO, error)
}

type TurnRepository struct {
	Store turnstore.TurnStoreInterface
}

func (tr *TurnRepository) Create(t domain.Turn) (*domain.Turn, error) {

	turn, err := tr.Store.Create(t)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (tr *TurnRepository) GetById(id int) (*domain.TurnDTO, error) {

	turn, err := tr.Store.GetById(id)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (tr *TurnRepository) FindById(id int) (*domain.Turn, error) {

	turn, err := tr.Store.FindById(id)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (tr *TurnRepository) GetAll() ([]domain.Turn, error) {

	turns, err := tr.Store.GetAll()
	if err != nil {
		return nil, err
	}

	return turns, nil

}

func (tr *TurnRepository) UpdateOne(id int, t domain.Turn) (*domain.Turn, error) {

	turn, err := tr.Store.UpdateOne(id, t)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (tr *TurnRepository) UpdateMany(id int, t domain.Turn) (*domain.Turn, error) {

	turn, err := tr.Store.UpdateMany(id, t)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (tr *TurnRepository) Delete(id int) error {

	if err := tr.Store.Delete(id); err != nil {
		return err
	}

	return nil

}

func (tr *TurnRepository) CreateByPatientDniAndDentistLicense(t domain.Turn, dni string, license string) (*domain.Turn, error) {

	turn, err := tr.Store.CreateByPatientDniAndDentistLicense(t, dni, license)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (tr *TurnRepository) GetByPatientDni(dni string) (*domain.TurnDTO, error) {

	turn, err := tr.Store.GetByPatientDni(dni)
	if err != nil {
		return nil, err
	}

	return turn, nil

}