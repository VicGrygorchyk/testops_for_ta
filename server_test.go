package main

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
	// jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func Test_bookmarkIndex(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/bookmark", bookmarkIndex())
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Test with std lib", func(t *testing.T) {
		res, err := http.Get(ts.URL + "/v1/bookmark")
		if err != nil {
		   t.Errorf("Expected nil, received %s", err.Error())
		}
		if res.StatusCode != http.StatusOK {
		   t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
		}
	})

	t.Run("test with apitest", func(t *testing.T) {
		apitest.New("Test bookmarkindex").
			Handler(r).
			Get("/v1/bookmark").
			Expect(t).
			Status(http.StatusOK).
			End()		
	})
 }

 func Test_bookmarkFind(t *testing.T) {
	r := mux.NewRouter()
	r.Handle("/v1/bookmark/{id}", bookmarkFind())
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("not found", func(t *testing.T) {
	   res, err := http.Get(ts.URL + "/v1/bookmark/1")
	   if err != nil {
		  t.Errorf("Expected nil, received %s", err.Error())
	   }
	   if res.StatusCode != http.StatusNotFound {
		  t.Errorf("Expected %d, received %d", http.StatusNotFound, res.StatusCode)
	   }
	})
	t.Run("found", func(t *testing.T) {
	   res, err := http.Get(ts.URL + "/v1/bookmark/2")
	   if err != nil {
		  t.Errorf("Expected nil, received %s", err.Error())
	   }
	   if res.StatusCode != http.StatusOK {
		  t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	   }
	})
 }
 