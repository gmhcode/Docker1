package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func connect() {
	//allows us to access the env file's values in os.GETEnv()
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file %v", err)
	}
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	d, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to open mysql %v", err)
	} else {
		fmt.Println("I am connected")
	}

	db = d
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Live change inside docker container. hello")
}

func main() {
	connect()
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8081", nil)

	fmt.Println("hello")
}
