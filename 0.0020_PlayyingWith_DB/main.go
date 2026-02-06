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
}

// One shall use migrations for actual use, this is just example
func createTable(db *sql.DB) {

	query := `create table Users
	id serial primary key,
	name varchar(20) not null,
	address varchar(50) not null
	`

	log.Println("Creating Table")
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

}
