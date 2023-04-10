package handlers

import (
	"net/http"
	"strconv"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/patient"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/web"
	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	PatientService patient.IPatientService
}

func (ph *PatientHandler) Post(c *gin.Context) {

	var patient domain.Patient

	err := c.ShouldBindJSON(&patient)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid JSON")
	}

	p, err := ph.PatientService.Create(patient)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
	}

	web.Success(c, http.StatusCreated, p)

}

func (ph *PatientHandler) GetById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	patient, err := ph.PatientService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "patient not found")
		return
	}

	web.Success(c, http.StatusOK, patient)

}

func (ph *PatientHandler) GetAll(c *gin.Context) {

	patients, err := ph.PatientService.GetAll()
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
	}

	web.Success(c, http.StatusOK, patients)

}

func (ph *PatientHandler) Patch(c * gin.Context){

	type Request struct {
		Name string
		Surname string
		Address string
		Dni string
		DischargeDate string
	}

	var r Request
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	update, err := ph.PatientService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "patient not found")
		return
	}

	err = c.ShouldBindJSON(&r)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid JSON")
		return
	}

	if r.Name != "" {
		update.Name = r.Name
	}

	if r.Surname != "" {
		update.Surname = r.Surname
	}

	if r.Address != "" {
		update.Address = r.Address
	}

	if r.Dni != "" {
		update.Dni = r.Dni
	}

	if r.DischargeDate != "" {
		update.DischargeDate = r.DischargeDate
	}

	result, err := ph.PatientService.UpdateOne(id, *update)
	if err != nil {
		web.NewApiError(c, http.StatusConflict,"conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, result)

}

func (ph *PatientHandler) Put (c *gin.Context) {

	var patient domain.Patient
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	_, err = ph.PatientService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "patient not found")
		return
	}

	err = c.ShouldBindJSON(&patient)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid JSON")
		return
	}

	p, err := ph.PatientService.UpdateMany(id, patient)
	if err != nil {
		web.NewApiError(c, http.StatusConflict,"conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, p)

}

func (ph *PatientHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	_, err = ph.PatientService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "patient not found")
		return
	}

	err = ph.PatientService.Delete(id)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, nil)

}
