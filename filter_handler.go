package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
)

func filterHandle(w http.ResponseWriter, r *http.Request) {
	//Handle 404 error
	if len(r.URL.Path) >= 7 {
		if r.URL.Path[0:8] != "/filter/" {
			sendError(w, 404)
			return
		}
	}
	path := r.URL.Path[8:]
	id := -1
	id, err := strconv.Atoi(r.URL.Path[8:])
	fmt.Println(path)
	if err != nil && path != "all" {
		sendError(w, 404)
		return
	}
	if (id <= 0 || id > len(All.Artists)) && path != "all" {
		sendError(w, 404)
		return
	}
	//Sending all artists
	if path == "all" {
		switch r.Method {
		case "GET":
			//Send all artists
			temp, err := template.ParseFiles("./static/templates/filter.html")
			if err != nil {
				sendError(w, 500)
				return
			}
			temp.Execute(w, All)
		case "POST":
			//Send filtered artists
			temp, err := template.ParseFiles("./static/templates/filter.html")
			if err != nil {
				sendError(w, 500)
				return
			}
			temp.Execute(w, filter(w, r))
		}

	} else {
		//Send one artist
		sendArtist(w, r, All.Artists[id-1].Name)
	}
}

func filter(w http.ResponseWriter, r *http.Request) API {
	//Main struct to return
	var ToSend API
	var ArtistsToSend Artists

	var Inputs struct {
		CreationDate struct {
			From string
			To   string
		}
		FirstAlbumDate struct {
			From string
			To   string
		}
		NumberOfMembers struct {
			From string
			To   string
		}
		Location string
	}

	//Check for invalid request data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendError(w, 400)
		return ToSend
	}
	query, err := url.ParseQuery(string(body))
	if err != nil {
		sendError(w, 400)
		return ToSend
	}
	for i, v := range query {
		switch i {
		case "cd-from":
			Inputs.CreationDate.From = v[0]
		case "cd-to":
			Inputs.CreationDate.To = v[0]
		case "fad-from":
			Inputs.FirstAlbumDate.From = v[0]
		case "fad-to":
			Inputs.FirstAlbumDate.To = v[0]
		case "nom-from":
			Inputs.NumberOfMembers.From = v[0]
		case "nom-to":
			Inputs.NumberOfMembers.To = v[0]
		case "loc":
			Inputs.Location = v[0]
		default:
			sendError(w, 400)
			return ToSend
		}
	}
	//Filtering by user inputs
	for i := 0; i < len(All.Artists); i++ {
		if !compareCreationDate(Inputs.CreationDate.From, Inputs.CreationDate.To, i) {
			continue
		}
		if !compareFirstAlbumDate(Inputs.FirstAlbumDate.From, Inputs.FirstAlbumDate.To, i) {
			continue
		}
		if !compareNumberOfMembers(Inputs.NumberOfMembers.From, Inputs.NumberOfMembers.To, i) {
			continue
		}
		if !compareLocation(Inputs.Location, i) {
			continue
		}
		ArtistsToSend = append(ArtistsToSend, All.Artists[i])
	}
	//return struct to send
	ToSend.Artists = ArtistsToSend
	return ToSend
}

//Compare users creation dates with Artists creation dates
func compareCreationDate(from string, to string, index int) bool {
	if from == "" && to == "" {
		return true
	}
	if from != "" && to != "" {
		compare := All.Artists[index].CreationDate
		fromN, err := strconv.Atoi(from)
		toN, err2 := strconv.Atoi(to)
		if err != nil || err2 != nil {
			return false
		}
		if compare >= fromN && compare <= toN {
			return true
		}
	}
	if from != "" && to == "" {
		compare := All.Artists[index].CreationDate
		fromN, err := strconv.Atoi(from)
		if err != nil {
			return false
		}
		if compare >= fromN {
			return true
		}
	}
	if from == "" && to != "" {
		compare := All.Artists[index].CreationDate
		toN, err := strconv.Atoi(to)
		if err != nil {
			return false
		}
		if compare <= toN {
			return true
		}
	}
	return false
}

//Compare users first album date with Artists first album date
func compareFirstAlbumDate(from string, to string, index int) bool {
	if from == "" && to == "" {
		return true
	}
	if from != "" && to != "" {
		fullDate := []rune(All.Artists[index].FirstAlbum)
		compare, err := strconv.Atoi(string(fullDate[len(fullDate)-4:]))
		fromN, err1 := strconv.Atoi(from)
		toN, err2 := strconv.Atoi(to)
		if err != nil || err1 != nil || err2 != nil {
			return false
		}
		if compare >= fromN && compare <= toN {
			return true
		}
	}
	if from != "" && to == "" {
		fullDate := All.Artists[index].FirstAlbum
		compare, err := strconv.Atoi(fullDate[len(fullDate)-4:])
		fromN, err := strconv.Atoi(from)
		if err != nil {
			return false
		}
		if compare >= fromN {
			return true
		}
	}
	if from == "" && to != "" {
		fullDate := All.Artists[index].FirstAlbum
		compare, err := strconv.Atoi(fullDate[len(fullDate)-4:])
		toN, err := strconv.Atoi(to)
		if err != nil {
			return false
		}
		if compare <= toN {
			return true
		}
	}
	return false
}

//Compare users number of members with Artists number of members
func compareNumberOfMembers(from string, to string, index int) bool {
	if from == "" && to == "" {
		return true
	}
	if from != "" && to != "" {
		compare := len(All.Artists[index].Members)
		fromN, err := strconv.Atoi(from)
		toN, err2 := strconv.Atoi(to)
		if err != nil || err2 != nil {
			return false
		}
		if compare >= fromN && compare <= toN {
			return true
		}
	}
	if from != "" && to == "" {
		compare := len(All.Artists[index].Members)
		fromN, err := strconv.Atoi(from)
		if err != nil {
			return false
		}
		if compare >= fromN {
			return true
		}
	}
	if from == "" && to != "" {
		compare := len(All.Artists[index].Members)
		toN, err := strconv.Atoi(to)
		if err != nil {
			return false
		}
		if compare <= toN {
			return true
		}
	}
	return false
}

//Compare users location with Artists locations
func compareLocation(location string, index int) bool {
	loc := strings.Split(location, ", ")
	city := strings.ToLower(loc[0])
	location = strings.ToLower(location)
	if location == "" {
		return true
	}

	for _, v := range All.Locations.Index[index].Locations {
		if strings.Index(strings.ToLower(v), location) != -1 {
			return true
		}
		if strings.Index(strings.ToLower(v), city) != -1 {
			return true
		}
	}

	return false
}
