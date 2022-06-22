package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

type Bookmark struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

func main() {
	router := mux.NewRouter()
	middleware := negroni.New(
		negroni.NewLogger(),
	)

	router.Handle(
		"/v1/bookmark", middleware.With(negroni.Wrap(bookmarkIndex())),
	).Methods("GET", "OPTIONS").Name("bookmarkIndex")

	router.Handle(
		"/v1/bookmark/{id}",
		middleware.With(negroni.Wrap(bookmarkFind())),
	).Methods("GET", "OPTIONS").Name("bookmarkFind")

	http.Handle("/", router)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	// server
	srv := &http.Server {
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr: ":8080",
		Handler: context.ClearHandler(http.DefaultServeMux),
		ErrorLog: logger,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func bookmarkIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := []*Bookmark{
			{
				ID:   1,
				Link: "http://google.com",
			},
			{
				ID:   2,
				Link: "https://apitest.dev",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading bookmarks"))
		}
	})
}

func bookmarkFind() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
		   w.WriteHeader(http.StatusInternalServerError)
		   w.Write([]byte("Error reading parameters"))
		   return
		}
		data := []*Bookmark{
		   {
			  ID:   2,
			  Link: "https://apitest.dev",
		   },
		}
		if id != data[0].ID {
		   w.WriteHeader(http.StatusNotFound)
		   w.Write([]byte("Not found"))
		   return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data[0]); err != nil {
		   w.WriteHeader(http.StatusInternalServerError)
		   w.Write([]byte("Error reading bookmark"))
		}
	 })
}
