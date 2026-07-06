package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	const dsn = "host=1.12.70.219 port=8596 user=postgres password=1234qazx@ dbname=cus_cms sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var username, oldHash string
	err = db.QueryRow("SELECT username, password_hash FROM bloggers LIMIT 1").Scan(&username, &oldHash)
	if err != nil {
		log.Fatal(err)
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte("test1234"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("UPDATE bloggers SET password_hash = $1", string(newHash))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("username=%s\nold_hash=%s\n", username, oldHash)
}
