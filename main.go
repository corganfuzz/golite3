package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//Book model
type Book struct {
	id     int
	name   string
	author string
}

func main() {
	fmt.Println("Server is running in port 8000...")

	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)

	if err != nil {
		log.Println(err)
	}

	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64), name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully create table books!")
	}
	statement.Exec()

	// Create
	statement, _ = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?,?,?)")
	statement.Exec("A tale of 2 cities", "Charles Dickens", 140430547)
	log.Println("Book Inserted into DB!")

	//Read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book

	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	// Update
	statement, _ = db.Prepare("update books set name=? where id=?")
	statement.Exec("The Tale of Two Cities, 1")
	log.Println("Successfully updated book in DB")

	//Delete

	statement, _ = db.Prepare("delete from books where id=?")
	statement.Exec(1)
	log.Println("Book deleted from DB!")

}
