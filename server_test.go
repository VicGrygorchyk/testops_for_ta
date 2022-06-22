package main

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
)

func Test_bookmarkIndex(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/bookmark", bookmarkIndex())
	r.Handle("/v1/bookmark/{id}", bookmarkFind())

	ts := httptest.NewServer(r)
	defer ts.Close()

	apitest.New("It should GET 200 ok for bookmark").
		Handler(r).
		Get("/v1/bookmark").
		Expect(t).
		Status(http.StatusOK).
		End()		
	
	apitest.New("It should find bookmark by index 1").
		Handler(r).
		Get("/v1/bookmark/1").
		Expect(t).
		Status(http.StatusOK).
		End()	

	apitest.New("It should find bookmark by index 2").
		Handler(r).
		Get("/v1/bookmark/2").
		Expect(t).
		Status(http.StatusOK).
		End()
 }
 