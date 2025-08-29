package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"TestGoProject/internal/api"
	"TestGoProject/internal/repository"
	"TestGoProject/internal/service"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "user:password@tcp(localhost:3306)/employeedb"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewEmployeeRepository(db)
	service := service.NewEmployeeService(repo)
	router := mux.NewRouter()
	api.RegisterEmployeeRoutes(router, service)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
