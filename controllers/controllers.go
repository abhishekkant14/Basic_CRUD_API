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
func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		rows, err := db.Query("SELECT id, name, email, age FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer rows.Close()

		var users []models.User

		for rows.Next() {

			var user models.User

			err := rows.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.Phone,
			)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}
		result, err := db.Exec(
			"UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?",
			user.Name,
			user.Email,
			user.Phone,
			id,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User updated successfully",
		})
	}
}
func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		result, err := db.Exec(
			"DELETE FROM users WHERE id = ?",
			id,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted successfully",
		})
	}
}
