package main

import (
	"bytes"
	"encoding/json"
	"musictracks/internal/models"
	"net/http"
	"net/url"
	"time"
)

// retreieve structured data from response body after GET request to spotify seacrh api
func getTrackDataFromResponseBody(m map[string]models.MusicTracks) (string, string, []string, string) {
	var items models.MusicTracks = m["tracks"]
	var popular, number int
	// if more than 1 items select the most popular one, with popularity value
	if len(items.Items) > 1 {
		for i := 0; i < len(items.Items); i++ {
			if items.Items[i].Popularity > popular {
				popular = items.Items[i].Popularity
				number = i
			}
		}
	}

	var item models.MusicTrack = items.Items[number]
	isrc := item.ExtID.ISRC
	title := item.Name
	imageURI := item.Album.Images[0].URL
	var artistsDetails []models.ArtistDetails = item.Artists
	var artists []string
	for _, v := range artistsDetails {
		artists = append(artists, v.Name)
	}
	return isrc, title, artists, imageURI
}

// get access token using client id and client secret
func getAccessToken() (string, error) {
	requestBody := url.Values{
		"grant_type":    {grantType},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
	}

	// create post request for retrieving access token
	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(requestBody.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// create a http client hit the URL
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", err
	}
	// retrieve the repsonse
	resBody := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&resBody)
	if err != nil {
		return "", err
	}

	return resBody["access_token"].(string), nil
}

// function to retrieve parameter value from url
func getValuefromRequest(vars map[string]string, value string) string {
	// check whether the isrc paramter is present
	v, ok := vars[value]
	if !ok || v == "" {
		return ""
	}
	return v
}

// retrieve Track data from track, image and artist models
func (app *application) getTrackDetails(isrc string) (string, string, []string, string) {
	title := app.tracks.GetTrackTitle(isrc)

	imageURI := app.images.GetImageURI(isrc)

	artistNames := app.artists.GetArtistNamesByISRC(isrc)

	return title, isrc, artistNames, imageURI

}
