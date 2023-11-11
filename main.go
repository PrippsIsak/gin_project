package main

import (
	"gin-twitter/app/routes"
	"gin-twitter/storage"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize your database connection
	dbPath := "db.db" // Provide the path to your SQLite database file
	storageInstance, err := storage.NewSqliteStorage(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer storageInstance.Close() // Defer closing the database connection

	// Define the table name and read the schema from the SQL file
	tableName := "persons"                                                           // Change this to your desired table name
	schemaFilePath := "/home/isak/goProjects/gin-twitter/storage/schemas/person.sql" // Path to your SQL schema file

	schemaBytes, readErr := os.ReadFile(schemaFilePath)
	if readErr != nil {
		log.Fatal(readErr)
	}

	schema := string(schemaBytes)

	// Call CreateTable to create the table with the specified schema
	err = storageInstance.CreateTable(tableName, schema)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	// Create a new Gin router
	router := gin.Default()

	// Initialize routes
	routes.InitRoutes(router)

	// Run the server
	router.Run(":8080")
}
