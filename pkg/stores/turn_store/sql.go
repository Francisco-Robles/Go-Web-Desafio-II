package turnstore

import (
	"database/sql"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
	dentiststore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/dentist_store"
	patientstore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/patient_store"
)

type TurnSqlStore struct {
	DB *sql.DB
	PatientStore patientstore.PatientStoreInterface
	DentistStore dentiststore.DentistStoreInterface
}

func (s *TurnSqlStore) Create(t domain.Turn) (*domain.Turn, error) {

	query := "INSERT INTO turns (description, patient_id, dentist_id) VALUES (?,?,?)"
	result, err := s.DB.Exec(query, t.Description, t.PatientId, t.DentistId)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	turn := &domain.Turn{
		Id: int(id),
		Description: t.Description,
		PatientId: t.PatientId,
		DentistId: t.DentistId,
	}

	return turn, nil

}

func (s *TurnSqlStore) GetById(id int) (*domain.Turn, error) {

	var turn domain.Turn

	query := "SELECT * FROM turns WHERE id = ?;"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&turn.Id, &turn.Description, &turn.PatientId, &turn.DentistId)
	if err != nil {
		return nil, err
	}

	return &turn, nil

}

func (s *TurnSqlStore) GetAll() ([]domain.Turn, error) {

	var turns []domain.Turn

	query := "SELECT * FROM turns;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var turn domain.Turn

		err := rows.Scan(&turn.Id, &turn.Description, &turn.PatientId, &turn.DentistId)
		if err != nil {
			return nil, err
		}

		turns = append(turns, turn)

	}

	return turns, nil

}

func (s *TurnSqlStore) UpdateOne(id int, t domain.Turn) (*domain.Turn, error) {

	query := "UPDATE turns SET description = ?, patient_id = ?, dentist_id = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, t.Description, t.PatientId, t.DentistId, id)
	if err != nil {
		return nil, err
	}

	turn := &domain.Turn{
		Id: id,
		Description: t.Description,
		PatientId: t.PatientId,
		DentistId: t.DentistId,
	}

	return turn, nil

}

func (s *TurnSqlStore) UpdateMany(id int, t domain.Turn) (*domain.Turn, error) {

	query := "UPDATE turns SET description = ?, patient_id = ?, dentist_id = ? WHERE id = ?;"
	_, err := s.DB.Exec(query, t.Description, t.PatientId, t.DentistId, id)
	if err != nil {
		return nil, err
	}

	turn := &domain.Turn{
		Id: id,
		Description: t.Description,
		PatientId: t.PatientId,
		DentistId: t.DentistId,
	}

	return turn, nil

}

func (s *TurnSqlStore) Delete(id int) error {

	query := "DELETE FROM turns WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}

func (s *TurnSqlStore) CreateByPatientDniAndDentistLicense(t domain.Turn, dni string, license string) (*domain.Turn, error) {

	idPatient, err := s.PatientStore.GetIdByDni(dni)
	if err != nil {
		return nil, err
	}

	idDentist, err := s.DentistStore.GetIdByLicense(license)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO turns (description, patient_id, dentist_id) VALUES (?,?,?);"
	result, err := s.DB.Exec(query, t.Description, idPatient, idDentist)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	turn := &domain.Turn{
		Id: int(id),
		Description: t.Description,
		PatientId: t.PatientId,
		DentistId: t.DentistId,
	}

	return turn, nil

}

func (s *TurnSqlStore) GetByPatientDni(dni string) (*domain.Turn, error) {

	var turn domain.Turn

	idPatient, err := s.PatientStore.GetIdByDni(dni)
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM turns WHERE patient_id = ?"
	row := s.DB.QueryRow(query, idPatient)

	err = row.Scan(&turn.Id, &turn.Description, &turn.PatientId, &turn.DentistId)
	if err != nil {
		return nil, err
	}

	return &turn, nil

}