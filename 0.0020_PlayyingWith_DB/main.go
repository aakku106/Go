package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:aakku106@localhost/gofirst?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	createTable(db)
	users := Users{
		Name:    "Aakku",
		Address: "Aakku_land 16",
	}
	pk := InsertIntoUserTable(db, users)
	log.Printf("ID: %d", pk)
	DisplayAllUserInfo(db)
}

// One shall use migrations for actual use, this is just example
func createTable(db *sql.DB) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		} // Althow this wont recover then we try to create same table more than 1 time cause we used log.fetal which uses os.exit

	}()

	query := `CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(20) NOT NULL,
	address VARCHAR(50) NOT NULL)
	`

	log.Println("Creating Table")
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
	log.Println(" Table Created")

}

type Users struct {
	Name    string
	Address string
}

// TO-DO create addData func to add data

func InsertIntoUserTable(db *sql.DB, user Users) int {
	query := `insert into Users(name,address)
	values($1,$2) returning id`
	var primaryKey int
	if err := db.QueryRow(query, user.Name, user.Address).Scan(&primaryKey); err != nil {
		log.Fatal(err)
	}
	return primaryKey
}
func DisplayAllUserInfo(db *sql.DB) {
	queue := `select * from Users`
	log.Println(db.Query(queue))
}
