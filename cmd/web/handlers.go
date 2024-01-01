package main

import (
	"encoding/json"
	"fmt"
	"musictracks/internal/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
)

// api method for creating track - GET /track/create/{ISRC}
func (app *application) createTrack(w http.ResponseWriter, r *http.Request) {
	// get access token
	accessToken, err := getAccessToken()
	serverError(w, err)

	// get value of isrc
	isrc := getValuefromRequest(mux.Vars(r), "ISRC")

	// return if track already exists
	trackTitle := app.tracks.GetTrackTitle(isrc)
	if trackTitle != "" {
		fmt.Fprintf(w, "Track: '%v' already exists", trackTitle)
		return
	}

	// add paramters to the url
	urlWithParams := searchURL + fmt.Sprintf("isrc:%s&type=%s", isrc, trackType)
	url, err := url.Parse(urlWithParams)
	serverError(w, err)

	// create a new GET request
	req, err := http.NewRequest("GET", url.String(), nil)
	serverError(w, err)

	fmt.Println("Write request: ", req)
	// add authorization token to header
	bearerToken := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Set("Authorization", bearerToken)

	// send a HTTP get request to search api
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	serverError(w, err)
	defer resp.Body.Close()

	// check if 200 status
	if resp.StatusCode != http.StatusOK {
		serverError(w, err)
	}

	// save response in map of struct
	var track = make(map[string]models.MusicTracks)
	err = json.NewDecoder(resp.Body).Decode(&track)
	serverError(w, err)

	// return if isrc not valid
	if len(track["tracks"].Items) == 0 {
		fmt.Fprint(w, "ISRC is nto valid")
		return
	}
	// get required data from response body
	isrc, title, artistnames, imageURI := getTrackDataFromResponseBody(track)

	// insert data into tables;
	err = app.tracks.Insert(isrc, title)
	serverError(w, err)
	err = app.artists.Insert(artistnames, isrc)
	serverError(w, err)
	err = app.images.Insert(imageURI, isrc)
	serverError(w, err)

	fmt.Fprintf(w, "Track Created-> Title: %v,  ISRC: %v,  Artist Names: %v, ImageURI: %v", title, isrc, artistnames, imageURI)

}

// api method to get single track GET /track/{ISRC}
func (app *application) getTrackByISRC(w http.ResponseWriter, r *http.Request) {
	isrc := getValuefromRequest(mux.Vars(r), "ISRC")
	title, isrc, artistNames, imageURI := app.getTrackDetails(isrc)
	fmt.Fprintf(w, "Track-> Title: %v,  ISRC: %v,  Artist Names: %v, ImageURI: %v", title, isrc, artistNames, imageURI)
}

// api method to get tracks using artist name - GET /artist/{name}
func (app *application) getTracksByArtist(w http.ResponseWriter, r *http.Request) {
	artistName := getValuefromRequest(mux.Vars(r), "name")
	isrcList := app.artists.GetISRCByName(artistName)
	for i := 0; i < len(isrcList); i++ {
		title, isrc, artistNames, imageURI := app.getTrackDetails(isrcList[i])
		fmt.Fprintf(w, "Track-> Title: %v,  ISRC: %v,  Artist Names: %v, ImageURI: %v \n", title, isrc, artistNames, imageURI)
	}
}
