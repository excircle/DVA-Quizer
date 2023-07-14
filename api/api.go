package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Question struct {
	ID            int    `json:"id"`
	Domain        string    `json:"domain"`
	Question      string `json:"question"`
	Answer        string `json:"answer"`
	Description   string `json:"description"`
	Miscellaneous bool   `json:"miscellaneous"`
}

func getDummy(w http.ResponseWriter, r *http.Request) {
    x := Question{ID: 1, Domain: "1", Question: "who?", Answer: "me!", Description: "desc", Miscellaneous: true}
    y := []Question{}
    y = append(y, x)
    json.NewEncoder(w).Encode(y)
}

func addQuestion(w http.ResponseWriter, r *http.Request) {
	log.Println("addQuestion is called")

	var question Question
	json.NewDecoder(r.Body).Decode(&question)

	// Open a connection to mysql
	db, err := sql.Open("mysql", "dva:sqlpass@tcp(172.18.0.4:3306)/dva")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create Transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Execute INSERT statement
	_, err = db.Exec("INSERT INTO dva.questions VALUES(NULL, ?, ?, ?, ?, ?)", question.Domain, question.Question, question.Answer, question.Description, question.Miscellaneous)
	if err != nil {
		log.Fatal(err)
	}

	// Close transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data:\n{\n\tDomain: %d,\n\tQuestion: '%s',\n\tAnswer: '%s'\n} Added Successfully", question.Domain, question.Question, question.Answer)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/questions", addQuestion).Methods("POST")
	r.HandleFunc("/test", getDummy).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://172.18.0.2:3000", "http://127.0.0.1:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions, http.MethodHead},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
