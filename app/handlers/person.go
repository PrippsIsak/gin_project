package handlers

import (
	"gin-twitter/storage"
	"gin-twitter/types"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPerson(c *gin.Context) {
	instance, err := storage.NewSqliteStorage("db.db")

	if err != nil {
		log.Println("Initilaztion err")
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	defer instance.Close()

	tableName := "persons"
	whereCondition := "username = ?"
	userName := c.Param("username")

	log.Println(userName)

	result, queryErr := instance.Get(tableName, whereCondition, userName)

	if queryErr != nil {
		log.Println("querry error!")
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "GetPerson succes",
		"result":  result,
	})
}

func CreatePerson(c *gin.Context) {
	var person types.Person

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
	}

	instance, err := storage.NewSqliteStorage("db.db")

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
	}

	defer instance.Close()

	if err := instance.Create("persons", &person); err != nil {
		c.JSON(500, gin.H{
			"error": "Error in creating person",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Person created successfully",
		"person":  person,
	})
}

func UpdatePerson(c *gin.Context) {
	var updatedFields map[string]interface{}

	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(400, gin.H{
			"error": "Bad request",
		})
		return
	}

	instance, err := storage.NewSqliteStorage("db.db")

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	defer instance.Close()

	userName := c.Param("username")
	whereCondition := "username = ?"
	setClasue := ""
	args := make([]interface{}, 0)

	for field, value := range updatedFields {
		setClasue += field + " = ?, "
		args = append(args, value)
	}

	setClasue = strings.TrimSuffix(setClasue, ", ")
	args = append(args, userName)

	if updateErr := instance.Update("persons", setClasue, whereCondition, args...); updateErr != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"Message": "succesfully updated" + userName,
		"fields":  updatedFields,
	})
}

func DeletePerson(c *gin.Context) {
	instance, err := storage.NewSqliteStorage("db.db")

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	defer instance.Close()

	tableName := "persons"
	whereCondition := "username = ?"
	userName := c.Param("username")

	if queryErr := instance.Delete(tableName, whereCondition, userName); queryErr != nil {
		c.JSON(500, gin.H{
			"error": "Server internal eroor",
		})
	}

	c.JSON(200, gin.H{
		"message": "Succes, user: " + userName + " was deleted",
	})
}
