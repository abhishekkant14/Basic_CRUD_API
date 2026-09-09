package main

import (
	"database/sql"
	"fmt"
	"log"

	"user_api/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:Abhishek@123456@tcp(127.0.0.1:3306)/user_db"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MySQL dB sucessfully Connected")

	router := gin.Default()

	router.POST("/users", controllers.CreateUser(db))

	router.Run(":8080")
}
