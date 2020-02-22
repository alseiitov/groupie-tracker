package main

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

//Send info about artist
func sendArtist(w http.ResponseWriter, r *http.Request, toSearch string) {
	All.ID = -1
	//Searching artist id by name
	for i := 0; i < 52; i++ {
		if strings.ToLower(All.Artists[i].Name) == strings.ToLower(toSearch) {
			All.ID = i
			break
		}
	}
	//Sending not found page if artist not found
	if All.ID == -1 {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		//Send info to user
		temp, err := template.ParseFiles("./static/templates/artist.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, All)
	}
}

//Send info about member
func sendMember(w http.ResponseWriter, r *http.Request, toSearch string) {
	type MemberPage struct {
		Title  string
		Artist []string
	}
	var memberPage MemberPage
	//Searching artist ids by member
	for i := 0; i < 52; i++ {
		for _, member := range All.Artists[i].Members {
			if strings.ToLower(member) == strings.ToLower(toSearch) {
				memberPage.Title = member + "<br>is a member of"
				memberPage.Artist = append(memberPage.Artist, All.Artists[i].Name)
			}
		}
	}
	//Sending not found page if artist not found
	if memberPage.Title == "" {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		//Sending info about member
		temp, err := template.ParseFiles("./static/templates/member.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, memberPage)
	}
}

//Send info about Location
func sendLocation(w http.ResponseWriter, r *http.Request, toSearch string) {
	type LocationPage struct {
		Title   string
		Artists []string
	}
	var locationPage LocationPage
	//Searching artist ids by city name
	for i, all := range All.Locations.Index {
		for _, location := range all.Locations {
			if strings.ToLower(location) == strings.ToLower(toSearch) {
				locationPage.Title = "Concerts in " + location
				locationPage.Artists = append(locationPage.Artists, All.Artists[i].Name)
				break
			}
		}
	}
	//Sending not found page if artist not found
	if locationPage.Title == "" {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		//Sending info about location
		temp, err := template.ParseFiles("./static/templates/location.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, locationPage)
	}
}

//Send info about first album
func sendFirstAlbum(w http.ResponseWriter, r *http.Request, toSearch string) {
	type FirstAlbumPage struct {
		Title   string
		Artists []string
	}
	var firstAlbumPage FirstAlbumPage
	//Searching artist ids by first album date
	for i, artist := range All.Artists {
		if strings.ToLower(artist.FirstAlbum) == strings.ToLower(toSearch) {
			firstAlbumPage.Title = "Artists / Bands relased their first album in " + artist.FirstAlbum
			firstAlbumPage.Artists = append(firstAlbumPage.Artists, All.Artists[i].Name)
		}
	}
	if firstAlbumPage.Title == "" {
		//Sending not found page if artist not found
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		//Sending info about location
		temp, err := template.ParseFiles("./static/templates/firstalbum.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, firstAlbumPage)
	}
}

//Send info about creation date
func sendCreationDate(w http.ResponseWriter, r *http.Request, toSearch string) {
	year, _ := strconv.Atoi(toSearch)
	type CreationDatePage struct {
		Title   string
		Artists []string
	}
	var creationDatePage CreationDatePage
	//Searching artist ids by first creation date
	for i, artist := range All.Artists {
		if artist.CreationDate == year {
			creationDatePage.Title = "Artists / Bands created in " + toSearch
			creationDatePage.Artists = append(creationDatePage.Artists, All.Artists[i].Name)
		}
	}
	if creationDatePage.Title == "" {
		//Sending not found page if artists not found
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		//Sending info about creation date
		temp, err := template.ParseFiles("./static/templates/creationdate.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, creationDatePage)
	}
}

//Send error
func sendError(w http.ResponseWriter, status int) {
	var title string
	switch status {
	case 400:
		title = "400 Bad Request"
	case 404:
		title = "404 Not Found"
	case 500:
		title = "500 Internal Sever Error"
	}

	w.WriteHeader(status)

	temp, err := template.ParseFiles("./static/templates/error.html")
	if err != nil {
		sendError(w, 500)
		return
	}
	temp.Execute(w, title)
	return
}
