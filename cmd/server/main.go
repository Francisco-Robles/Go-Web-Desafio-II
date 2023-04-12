package main

import (
	"database/sql"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/cmd/server/handlers"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/dentist"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/patient"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/turn"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/middleware"
	dentiststore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/dentist_store"
	patientstore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/patient_store"
	turnstore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/turn_store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load("./.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

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

	turnStorage := turnstore.TurnSqlStore{DB: db, PatientStore: &patientStorage, DentistStore: &dentistStorage }
	turnRepo := turn.TurnRepository{Store: &turnStorage}
	turnService := turn.TurnService{Repository: &turnRepo}
	turnHandler := handlers.TurnHandler{TurnService: &turnService}
	
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Authentication(), middleware.Logger())

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

	turns := r.Group("turns")
	{
		turns.POST("", turnHandler.Post)
		turns.GET("", turnHandler.GetAll)
		turns.GET(":id", turnHandler.GetById)
		turns.PATCH(":id", turnHandler.Patch)
		turns.PUT(":id", turnHandler.Put)
		turns.DELETE(":id", turnHandler.Delete)
		turns.POST("/byDniAndLicense", turnHandler.PostByPatientDniAndDentistLicense)
		turns.GET("/byDniPatient", turnHandler.GetByPatientDni)
	}

	r.Run()

}