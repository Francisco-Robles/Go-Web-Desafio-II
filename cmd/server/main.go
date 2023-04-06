package main

import (
	"database/sql"

	"github.com/Francisco-Robles/Go-Web-Desafio-II/cmd/server/handlers"
	"github.com/Francisco-Robles/Go-Web-Desafio-II/internal/layers/dentist"
	dentiststore "github.com/Francisco-Robles/Go-Web-Desafio-II/pkg/stores/dentist_store"
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

	dentistStorage := dentiststore.DentistSqlStore{db}
	dentistRepo := dentist.DentistRepository{&dentistStorage}
	dentistService := dentist.DentistService{&dentistRepo}
	dentistHandler := handlers.DentistHandler{&dentistService}

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

	r.Run()

}