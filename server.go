package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-student-api/graph"
	"github.com/go-student-api/graph/generated"
	"github.com/go-student-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


const defaultPort = "8080"
var db *gorm.DB
func initDB(){

	var err error
	dataSourceName := "host=localhost port=5432 user=postgres dbname=school_modules sslmode=disable password=admin"
	db, err = gorm.Open("postgres", dataSourceName)
	if err != nil {
	   fmt.Println(err)
	   panic("Failed to connect to database")
	}

	db.LogMode(true)
	 // Create the database. This is a one-time step.
    // Comment out if running multiple times - You may see an error otherwise
	// db.Exec("CREATE DATABASE student_modular_db")
	// db.Exec("USE student_modular_db")

	 // Migration to create tables
	db.AutoMigrate(&models.Student{})

	
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
