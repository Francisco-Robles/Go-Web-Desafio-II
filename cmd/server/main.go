package main

import (
	"database/sql"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/cmd/server/handlers"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/dentist"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/patient"
	dentiststore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/dentist_store"
	patientstore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/patient_store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:elescuadron10@tcp(localhost:3306)/desafio_II")
	if err != nil {
		panic(err)
	}

	errPing := db.Ping()
	if err != nil {
		panic(errPing)
	}

	dentistStorage := dentiststore.DentistSqlStore{DB: db}
	dentistRepo := dentist.DentistRepository{Store: &dentistStorage}
	dentistService := dentist.DentistService{Repository: &dentistRepo}
	dentistHandler := handlers.DentistHandler{DentistService: &dentistService}

	patientStorage := patientstore.PatientSqlStore{DB: db}
	patientRepo := patient.PatientRepository{Store: &patientStorage}
	patientService := patient.PatientService{Repository: &patientRepo}
	patientHandler := handlers.PatientHandler{PatientService: &patientService}

	r := gin.Default()

	dentists := r.Group("dentists")
	{
		dentists.POST("", dentistHandler.Post)
		dentists.GET("", dentistHandler.GetAll)
		dentists.GET(":id", dentistHandler.GetById)
		dentists.PATCH(":id", dentistHandler.Patch)
		dentists.PUT(":id", dentistHandler.Put)
		dentists.DELETE(":id", dentistHandler.Delete)

	}

	patients := r.Group("patients")
	{
		patients.POST("", patientHandler.Post)
		patients.GET("", patientHandler.GetAll)
		patients.GET(":id", patientHandler.GetById)
		patients.PATCH(":id", patientHandler.Patch)
		patients.PUT(":id", patientHandler.Put)
		patients.DELETE(":id", patientHandler.Delete)
	}

	r.Run()

}