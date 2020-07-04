package database

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var dbDriver,
	dbUser,
	dbPass,
	dbName,
	dbHost,
	dbPort string

func configDataBase() {
	file, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var readResults []string
	for scanner.Scan() {
		configData := strings.Split(scanner.Text(), "=")[1]
		readResults = append(readResults, configData)
	}
	dbDriver = readResults[0]
	dbUser = readResults[1]
	dbPass = readResults[2]
	dbHost = readResults[3]
	dbPort = readResults[4]
	dbName = readResults[5]
}

// Connect ke Database
func ConnectDB() (db *sql.DB) {
	configDataBase()
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Print(err)
		fmt.Scanln()
		log.Fatal(err)
	}
	fmt.Println("DataBase Successfully Connected")
	return db
}
