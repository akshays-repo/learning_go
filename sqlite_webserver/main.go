package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func main() {

	ConnectDatabase()

	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("person", getPersons)
		v1.GET("person/:id", getPersonById)
		v1.POST("person", addPerson)
		v1.PUT("person/:id", updatePerson)
		v1.DELETE("person/:id", deletePerson)
		v1.OPTIONS("person", options)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.

	r.Run()

}

func ConnectDatabase() error {

	db, err := sql.Open("sqlite3", "./names.db")

	if err != nil {
		return err
	}
	DB = db

	return nil
}

func getPersons(c *gin.Context) {

	persons, error := getPersonsRepo(DB, 9000)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}

	c.JSON(http.StatusOK, gin.H{"persons": persons})
}

func getPersonById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "getPersonById " + id + " Called"})
}

func addPerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "addPerson Called"})
}

func updatePerson(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "updatePerson Called"})
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "deletePerson " + id + " Called"})
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "options Called"})
}
