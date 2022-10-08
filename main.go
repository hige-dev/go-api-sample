package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=localhost user=root dbname=testdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/users", Users)

	server := &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Printf("error: %+v", err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func Users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetUsers(w, r)
	case "POST":
		// user作成
		// jsonをパースしてDBに保存
		//
	default:
		fmt.Fprintf(w, "Unauthorized Method: %+v", r.Method)
	}
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	rows, err := Db.Query("SELECT id, name, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}

		if err = rows.Scan(&user.Id, &user.Name, &user.CreatedAt); err != nil {
			fmt.Fprint(w, err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		log.Print("user is not exists.")
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&users); err != nil {
		// fmt.Fprint(w, err)
		log.Fatal(err)
	}
	if _, err := fmt.Fprintf(w, buf.String()); err != nil {
		fmt.Fprint(w, err)
	}
	rows.Close()
}
