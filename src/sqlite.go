// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/mattn/go-sqlite3"
// 	"log"
// )


// type Task interface {
//     create()
//     get()
//     edit()
//     delete()
// }
// type Tasks struct {
// 	task string
// 	id   int
// }
// func main() {
// 	db, err := sql.Open("sqlite3", "task.db")
// 	if err != nil {
// 		log.Fatal(err)
// 		_, err = db.Exec("CREATE TABLE tasks(id INTEGER PRIMARY KEY, task TEXT);")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	defer db.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("table tasks created")
// }
