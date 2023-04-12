package turn

import (

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
)

type ITurnService interface {
	Create(t domain.Turn) (*domain.Turn, error)
	GetById(id int) (*domain.TurnDTO, error)
	FindById(id int) (*domain.Turn, error)
	GetAll() ([]domain.Turn, error)
	UpdateOne(id int, t domain.Turn) (*domain.Turn, error)
	UpdateMany(id int, t domain.Turn) (*domain.Turn, error)
	Delete(id int) error
	CreateByPatientDniAndDentistLicense(t domain.Turn, dni string, license string) (*domain.Turn, error)
	GetByPatientDni(dni string) (*domain.TurnDTO, error)
}

type TurnService struct {
	Repository ITurnRepository
}

func NewTurnService(r ITurnRepository) ITurnService {
	return &TurnService{Repository: r}
}

func (ts *TurnService) Create(t domain.Turn) (*domain.Turn, error) {

	turn, err := ts.Repository.Create(t)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (ts *TurnService) GetById(id int) (*domain.TurnDTO, error) {

	turn, err := ts.Repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (ts *TurnService) FindById(id int) (*domain.Turn, error) {

	turn, err := ts.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (ts *TurnService) GetAll() ([]domain.Turn, error) {

	turns, err := ts.Repository.GetAll()
	if err != nil {
		return nil, err
	}

	return turns, nil

}

func (ts *TurnService) UpdateOne(id int, t domain.Turn) (*domain.Turn, error) {

	turn, err := ts.Repository.UpdateOne(id, t)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (ts *TurnService) UpdateMany(id int, t domain.Turn) (*domain.Turn, error) {

	turn, err := ts.Repository.UpdateMany(id, t)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (ts *TurnService) Delete(id int) error {

	if err := ts.Repository.Delete(id); err != nil {
		return err
	}

	return nil

}

func (ts *TurnService) CreateByPatientDniAndDentistLicense(t domain.Turn, dni string, license string) (*domain.Turn, error) {

	turn, err := ts.Repository.CreateByPatientDniAndDentistLicense(t, dni, license)
	if err != nil {
		return nil, err
	}

	return turn, nil

}

func (ts *TurnService) GetByPatientDni(dni string) (*domain.TurnDTO, error) {

	turn, err := ts.Repository.GetByPatientDni(dni)
	if err != nil {
		return nil, err
	}

	return turn, nil

}