package controllers

import (
	"database/sql"
	"net/http"

	"user_api/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User

		// Read JSON body
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Insert data into MySQL
		result, err := db.Exec(
			"INSERT INTO users (name, email, age) VALUES (?, ?, ?)",
			user.Name,
			user.Email,
			user.Phone,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		user.ID = int(id)

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    user,
		})
	}
}
