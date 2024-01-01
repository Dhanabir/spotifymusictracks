package main

import "net/http"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func serverError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		panic(err)
	}
}

func badRequestError(w http.ResponseWriter, v string) {
	if v == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
