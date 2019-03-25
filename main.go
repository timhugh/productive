package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbFile := flag.String("db", "visits.db", "Specify the path of a sqlite database to use.")
	db, err := connectDB(*dbFile)
	checkFatal(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()

		// is this safe for concurrency? :thinking_face:
		if _, err := db.Exec("INSERT INTO visits(url, timestamp) VALUES (?, julianday('now'))", url); err != nil {
			fmt.Printf("Error while inserting visit for %s: %v\n", url, err)
		}

		var count int
		row := db.QueryRow("SELECT COUNT(timestamp) FROM visits WHERE url = ?", url)
		err := row.Scan(&count)
		if err != nil {
			fmt.Printf("Error while counting visits for %s: %v\n", url, err)

			fmt.Fprint(w, "Blocked!")
			return
		}

		fmt.Printf("Logged visit %d to %s\n", count, url)
		fmt.Fprintf(w, "Blocked! You've attempted to visit %d times.\n", count)
	})

	forever := make(chan struct{})
	go func() {
		checkFatal(http.ListenAndServe(":80", nil))
	}()
	go func() {
		checkFatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
	}()
	<-forever
}

func connectDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS visits (url TEXT, timestamp REAL)")
	return db, err
}

func checkFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
