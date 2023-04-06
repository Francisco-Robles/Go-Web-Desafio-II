package handlers

import (
	"net/http"
	"strconv"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/domain"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/dentist"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/web"
	"github.com/gin-gonic/gin"
)

type DentistHandler struct {
	DentistService dentist.IDentistService
}

func (dh *DentistHandler) Post(c *gin.Context) {

	var dentist domain.Dentist

	err := c.ShouldBindJSON(&dentist)
	if err != nil {
		web.NewBadRequestApiError("invalid JSON")
	}

	d, err := dh.DentistService.Create(dentist)
	if err != nil {
		web.NewBadRequestApiError(err.Error())
	}

	web.Success(c, http.StatusCreated, d)

}

func (dh *DentistHandler) GetById(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewBadRequestApiError("invalid id.")
		return
	}

	dentist, err := dh.DentistService.GetById(id)
	if err != nil {
		web.NewBadRequestApiError("dentist not found.")
		return
	}

	web.Success(c, http.StatusOK, dentist)

}

func (dh *DentistHandler) GetAll(c *gin.Context) {

	dentists, err := dh.DentistService.GetAll()
	if err != nil {
		web.NewBadRequestApiError(err.Error())
	}

	web.Success(c, http.StatusOK, dentists)

}

func (dh *DentistHandler) Patch(c *gin.Context) {

	type Request struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		License string `json:"license"`
	}

	var r Request
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewBadRequestApiError("invalid id.")
		return
	}

	dentistBack, err := dh.DentistService.GetById(id)
	if err != nil {
		web.NewNotFoundApiError("dentist not found.")
		return
	}

	if err := c.ShouldBindJSON(&r); err != nil {
		web.NewBadRequestApiError("invalid JSON.")
		return
	}

	var update domain.Dentist

	if r.Name == "" {
		if r.Surname == "" {
			update = domain.Dentist{
				Name:    dentistBack.Name,
				Surname: dentistBack.Surname,
				License: r.License,
			}
		}else{
			update = domain.Dentist{
				Name:    dentistBack.Name,
				Surname: r.Surname,
				License: dentistBack.License,
			}
		}

	}else{
		update = domain.Dentist{
			Name:    r.Name,
			Surname: dentistBack.Surname,
			License: dentistBack.License,
		}
	}

	d, err := dh.DentistService.UpdateOne(id, update)
	if err != nil {
		web.NewConflictApiError(err.Error())
		return
	}

	web.Success(c, http.StatusOK, d)

}

func (dh *DentistHandler) Put(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewBadRequestApiError("invalid id.")
		return
	}

	_, err = dh.DentistService.GetById(id)
	if err != nil {
		web.NewNotFoundApiError("dentist not found.")
		return
	}

	var dentist domain.Dentist
	err = c.ShouldBindJSON(&dentist)
	if err != nil {
		web.NewBadRequestApiError("invalid JSON.")
		return
	}

	d, err := dh.DentistService.UpdateMany(id, dentist)
	if err != nil {
		web.NewConflictApiError(err.Error())
		return
	}

	web.Success(c, http.StatusOK, d)

}

func (dh *DentistHandler) Delete(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		web.NewBadRequestApiError("invalid id.")
		return
	}

	err = dh.DentistService.Delete(id)
	if err != nil {
		web.NewNotFoundApiError("dentist not found.")
	}

	web.Success(c, http.StatusOK, nil)

}
