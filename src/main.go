package main

import (
	"fmt"
	//"net/http"
	"database/sql"
	//"log"
	//"pq"
	_ "github.com/lib/pq"
)

var pool *sql.DB

const (
	DB_USER     = "postgres"
	//DB_PASSWORD = ""
	DB_NAME     = "customcommit"
)

func checkErr(err error) {
	if err != nil {
	  fmt.Println(err)
	}
  }

func connectToDB () {
	dbinfo := fmt.Sprintf("user=%s host=/tmp dbname=%s sslmode=disable",
            DB_USER, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
	fmt.Println("DB connection")
    defer db.Close()
	fmt.Println("# Querying")
        rows, err := db.Query("SELECT * FROM commits")
        checkErr(err)

        for rows.Next() {
            var cId int
			var cType int
            var commitText string
            err = rows.Scan(&cId, &cType, &commitText)
            checkErr(err)
            fmt.Println("cId | cType | commitText ")
            fmt.Printf("%3v | %8v | %6v\n", cId, cType, commitText)
        }
}

// func getGachiPhrase(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Fucking cumming")
// }

func main() {
	//http.HandleFunc("/", getGachiPhrase)
    //http.ListenAndServe(":8080", nil)
	connectToDB()
}