package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

var persons = []person{
	{ID: "1", Name: "John Doe", Age: 25, City: "New York"},
	{ID: "2", Name: "Alice Smith", Age: 30, City: "San Francisco"},
	{ID: "3", Name: "Bob Johnson", Age: 28, City: "Los Angeles"},
}

func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func personById(c *gin.Context) {
	id := c.Param("id")
	person, err := getPersonById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, person)
}

func createPerson(c *gin.Context) {
	var newPerson person

	if err := c.BindJSON(&newPerson); err != nil {
		return
	}

	persons = append(persons, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}

func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.GET("/persons/:id", personById)
	router.POST("/persons", createPerson)
	router.Run("localhost:8080")
}

func getPersonById(id string) (*person, error) {
	for i, p := range persons {
		if p.ID == id {
			return &persons[i], nil
		}
	}

	return nil, errors.New("person not found")
}
