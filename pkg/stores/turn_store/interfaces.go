package turnstore

import "github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"

type TurnStoreInterface interface {
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
