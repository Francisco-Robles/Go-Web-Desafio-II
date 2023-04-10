package dentiststore

import (
	"database/sql"
	"fmt"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
)

type DentistSqlStore struct {
	DB *sql.DB
}

func (s *DentistSqlStore) Create(d domain.Dentist) (*domain.Dentist, error) {

	var dentist domain.Dentist

	query := "INSERT INTO dentists (name, surname, license) VALUES (?,?,?);"
	row := s.DB.QueryRow(query, d.Name, d.Surname, d.License)

	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.Surname, &dentist.License)
	if err != nil {
		return nil, err
	}

	return &dentist, nil

}

func (s *DentistSqlStore) GetById(id int) (*domain.Dentist, error) {

	var dentist domain.Dentist

	query := "SELECT * FROM dentists WHERE id = ?;"
	row := s.DB.QueryRow(query, id)

	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.Surname, &dentist.License)
	if err != nil {
		return nil, err
	}

	return &dentist, nil

}
func (s *DentistSqlStore) GetAll() ([]domain.Dentist, error) {

	var dentists []domain.Dentist

	query := "SELECT * FROM dentists;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var dentist domain.Dentist

		err := rows.Scan(&dentist.Id, &dentist.Name, &dentist.Surname, &dentist.License)
		if err != nil{
			return nil, err
		}

		dentists = append(dentists, dentist)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dentists, nil

}

func (s *DentistSqlStore) UpdateOne(id int, d domain.Dentist) (*domain.Dentist, error) {

	var dentist domain.Dentist

	query := "UPDATE dentists SET name = ?, surname = ?, license = ? WHERE id = ?;"
	row := s.DB.QueryRow(query, d.Name, d.Surname, d.License, id)

	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.Surname, &dentist.License)
	if err != nil {
		return nil, err
	}

	return &dentist, nil

}

func (s *DentistSqlStore) UpdateMany(id int, d domain.Dentist) (*domain.Dentist, error) {

	var dentist domain.Dentist

	query := "UPDATE dentists SET name = ?, surname = ?, license = ? WHERE id = ?;"
	row := s.DB.QueryRow(query, d.Name, d.Surname, d.License, id)

	err := row.Scan(&dentist.Id, &dentist.Name, &dentist.Surname, &dentist.License)
	if err != nil {
		return nil, err
	}

	return &dentist, nil

}

func (s *DentistSqlStore) Delete(id int) error {

	query := "DELETE FROM dentists WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return nil

}

func (s *DentistSqlStore) GetIdByLicense(license string) (int, error) {

	var id int

	query := "SELECT id FROM dentists WHERE license = ?"
	row := s.DB.QueryRow(query, license)

	err := row.Scan(&id)
	fmt.Println(err)
	if err != nil {
		return -1, err
	}

	return id, nil

}