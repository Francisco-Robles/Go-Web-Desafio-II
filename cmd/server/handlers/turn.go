package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/turn"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/web"
	"github.com/gin-gonic/gin"
)

type TurnHandler struct {
	TurnService turn.ITurnService
}

func NewTurnHandler(s turn.ITurnService) *TurnHandler {
	return &TurnHandler{TurnService: s}
}

func (th *TurnHandler) Post(c * gin.Context) {

	var turn domain.Turn

	err := c.ShouldBindJSON(&turn)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid JSON")
		return
	}

	t, err := th.TurnService.Create(turn)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusCreated, t)

}

func (th *TurnHandler) GetById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	turn, err := th.TurnService.GetById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "turn not found")
		return
	}

	web.Success(c, http.StatusOK, turn)

}

func (th *TurnHandler) GetAll(c *gin.Context) {

	turns, err := th.TurnService.GetAll()
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, turns)

}

func (th *TurnHandler) Patch(c *gin.Context) {

	type Request struct {
		Datetime string `json:"datetime"`
		Description string `json:"description"`
		PatientId int `json:"patient_id"`
		DentistId int `json:"dentist_id"`
	}

	var r Request
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	update, err := th.TurnService.FindById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "turn not found")
		return
	}

	if err := c.ShouldBindJSON(&r); err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid JSON")
		return
	}

	if r.Datetime != "" {
		update.DateTime = r.Datetime
	}

	if r.Description != "" {
		update.Description = r.Description
	}

	if r.PatientId != 0 {
		update.PatientId = r.PatientId
	}

	if r.DentistId != 0 {
		update.DentistId = r.DentistId
	}

	t, err := th.TurnService.UpdateOne(id, *update)
	if err != nil {
		web.NewApiError(c, http.StatusConflict,"conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, t)

}

func (th *TurnHandler) Put(c *gin.Context) {

	var turn domain.Turn
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	_, err = th.TurnService.FindById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "turn not found")
		return
	}

	err = c.ShouldBindJSON(&turn)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid JSON")
		return
	}

	t, err := th.TurnService.UpdateMany(id, turn)
	if err != nil {
		web.NewApiError(c, http.StatusConflict,"conflict", err.Error())
		return
	}

	web.Success(c, http.StatusOK, t)

}

func (th *TurnHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid id")
		return
	}

	_, err = th.TurnService.FindById(id)
	if err != nil {
		web.NewApiError(c, http.StatusNotFound,"not_found", "turn not found")
		return
	}

	err = th.TurnService.Delete(id)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, fmt.Sprintf("Turn id = %d has been deleted.", id))

}

func (th *TurnHandler) PostByPatientDniAndDentistLicense(c *gin.Context) {

	dni := c.Query("dni")
	license := c.Query("license")
	var turn domain.Turn
	err := c.ShouldBindJSON(&turn)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", "invalid JSON")
		return
	}

	t, err := th.TurnService.CreateByPatientDniAndDentistLicense(turn, dni, license)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
		return
	}
	web.Success(c, http.StatusCreated, t)

}

func (th *TurnHandler) GetByPatientDni(c *gin.Context) {

	dni := c.Query("dni")

	turn, err := th.TurnService.GetByPatientDni(dni)
	if err != nil {
		web.NewApiError(c, http.StatusBadRequest,"bad_request", err.Error())
		return
	}

	web.Success(c, http.StatusOK, turn)

}
