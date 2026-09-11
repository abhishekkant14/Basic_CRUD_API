package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"user_api/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// MySQL connection
	dsn := "root:Abhishek@123456@tcp(127.0.0.1:3306)/user_db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check MySQL connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MySQL dB successfully Connected")

	// Create Gin router
	router := gin.Default()

	// CREATE user
	router.POST("/users", controllers.CreateUser(db))

	// GET all users
	router.GET("/users", controllers.GetUsers(db))

	router.PUT("/users/:id", controllers.UpdateUser(db))

	// Start server
	router.Run(":8080")
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
