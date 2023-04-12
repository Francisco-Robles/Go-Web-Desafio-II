package patientstore

import (
	"database/sql"
	"fmt"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
)

type PatientSqlStore struct {
	DB *sql.DB
}

func NewPatientSqlStore (db *sql.DB) PatientStoreInterface {
	return &PatientSqlStore{DB: db}
}

func (s *PatientSqlStore) Create(p domain.Patient) (*domain.Patient, error) {

	var patient domain.Patient
	
	query := "INSERT INTO patients (name, surname, address, dni, discharge_date) VALUES (?,?,?,?,?)"
	row := s.DB.QueryRow(query, p.Name, p.Surname, p.Address, p.Dni, p.DischargeDate)
	
	err := row.Scan(&patient.Id, &patient.Name, &patient.Surname, &patient.Address, &patient.Dni, &patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	return &patient, nil

}

func (s *PatientSqlStore) GetById(id int) (*domain.Patient, error) {

	var patient domain.Patient

	query := "SELECT * FROM patients WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&patient.Id, &patient.Name, &patient.Surname, &patient.Address, &patient.Dni, &patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	return &patient, nil

}

func (s *PatientSqlStore) GetAll() ([]domain.Patient, error) {

	var patients []domain.Patient

	query := "SELECT * FROM patients"
	rows, err := s.DB.Query(query)
	if err != nil {
		fmt.Println("error 1")
		return nil, err
	}
	
	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient

		err := rows.Scan(&patient.Id, &patient.Name, &patient.Surname, &patient.Address, &patient.Dni, &patient.DischargeDate)
		if err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil

}

func (s *PatientSqlStore) UpdateOne(id int, p domain.Patient) (*domain.Patient, error) {

	var patient domain.Patient

	query := "UPDATE patients SET name = ?, surname = ?, address = ?, dni = ?, discharge_date = ? WHERE id = ?"
	row := s.DB.QueryRow(query, p.Name, p.Surname, p.Address, p.Dni, p.DischargeDate, id)

	err := row.Scan(&patient.Id, &patient.Name, &patient.Surname, &patient.Address, &patient.Dni, & patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	return &patient, nil

}

func (s *PatientSqlStore) UpdateMany(id int, p domain.Patient) (*domain.Patient, error) {

	var patient domain.Patient

	query := "UPDATE patients SET name = ?, surname = ?, address = ?, dni = ?, discharge_date = ? WHERE id = ?"
	row := s.DB.QueryRow(query, p.Name, p.Surname, p.Address, p.Dni, p.DischargeDate, id)

	err := row.Scan(&patient.Id, &patient.Name, &patient.Surname, &patient.Address, &patient.Dni, & patient.DischargeDate)
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (s *PatientSqlStore) Delete(id int) error {

	query := "DELETE FROM patients WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil

}

func (s *PatientSqlStore) GetIdByDni (dni string) (int, error) {

	var id int

	query := "SELECT id FROM patients WHERE dni = ?"
	result := s.DB.QueryRow(query, dni)

	err := result.Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil

}