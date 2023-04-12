DROP DATABASE IF EXISTS `desafio_II`;

CREATE DATABASE `desafio_II`;
USE `desafio_II`;

CREATE TABLE `dentists` (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(55),
    surname VARCHAR(55),
    license VARCHAR(55)
);

INSERT INTO dentists (name, surname, license) VALUES ("Lionel", "Messi", "ABDC123");
INSERT INTO dentists (name, surname, license) VALUES ("Ángel", "Di María", "NSAS654");
INSERT INTO dentists (name, surname, license) VALUES ("Julián", "Álvarez", "HBSA789");
INSERT INTO dentists (name, surname, license) VALUES ("Lisandro", "Martínez", "ABC123");

CREATE TABLE `patients` (
	id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(55),
    surname VARCHAR(55),
    address VARCHAR(55),
    dni VARCHAR(20),
    discharge_date VARCHAR(15)
);

INSERT INTO patients (name, surname, address, dni, discharge_date) VALUES ("Facundo", "Conte", "ABDC123", "35456122", "2023-08-23");
INSERT INTO patients (name, surname, address, dni, discharge_date) VALUES ("Bruno", "Lima", "JNF415", "35156748", "2023-05-14");
INSERT INTO patients (name, surname, address, dni, discharge_date) VALUES ("Luciano", "De Cecco", "DJA456", "35689548", "2023-07-27");
INSERT INTO patients (name, surname, address, dni, discharge_date) VALUES ("Agustín", "Loser", "DAS456", "39658452", "2023-11-11");

CREATE TABLE `turns`(
	id INT PRIMARY KEY AUTO_INCREMENT,
    datetime VARCHAR(25),
    description VARCHAR(255),
    patient_id INT,
    dentist_id INT,
    CONSTRAINT fk_patient FOREIGN KEY (patient_id) REFERENCES patients(id),
    CONSTRAINT fk_dentist FOREIGN KEY (dentist_id) REFERENCES dentists(id)
);

INSERT INTO turns (datetime, description, patient_id, dentist_id) VALUES ("2023-04-11 15:30:00", "Limpieza y revisión dental", 2, 1);
INSERT INTO turns (datetime, description, patient_id, dentist_id) VALUES ("2023-04-10 16:30:00", "Extracción de muela del juicio", 3, 2);
INSERT INTO turns (datetime, description, patient_id, dentist_id) VALUES ("2023-04-09 17:30:00", "Tratamiento de conducto en diente #28", 2, 3);
INSERT INTO turns (datetime, description, patient_id, dentist_id) VALUES ("2023-04-08 18:00:00", "Consulta para ortodoncia", 1, 4);

CREATE VIEW get_turnDTO AS
SELECT t.id, t.datetime, t.description, p.name AS "patient_name", p.surname AS "patient_surname", p.address, p.dni, d.name AS "dentist_name", d.surname AS "dentist_surname", d.license
FROM turns t
INNER JOIN patients p ON t.patient_id = p.id
INNER JOIN dentists d ON t.dentist_id = d.id;

