package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/track/{ISRC}", app.getTrackByISRC).Methods("GET")
	r.HandleFunc("/track/create/{ISRC}", app.createTrack).Methods("GET")
	r.HandleFunc("/artist/{name}", app.getTracksByArtist).Methods("GET")
	return r
}
