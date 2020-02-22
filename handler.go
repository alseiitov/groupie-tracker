package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"
)

func indexHandle(w http.ResponseWriter, r *http.Request) {
	//Handle 404 error
	if r.URL.Path != "/" {
		sendError(w, 404)
		return
	}
	//Check for request method
	switch r.Method {
	case "GET":
		temp, err := template.ParseFiles("./static/templates/index.html")
		if err != nil {
			sendError(w, 500)
			return
		}
		temp.Execute(w, All)
	case "POST":
		var toSearch string   //users input
		var searchType string //what to search

		//Check for invalid request data
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			sendError(w, 400)
			return
		}
		query, err := url.ParseQuery(string(body))
		if err != nil {
			sendError(w, 400)
			return
		}
		for i, v := range query {
			switch i {
			case "toSearch":
				toSearch = v[0]
			case "searchType":
				searchType = v[0]
			default:
				sendError(w, 400)
				return
			}
		}
		if searchType != "artist" && searchType != "member" && searchType != "location" && searchType != "firstAlbum" && searchType != "creationDate" {
			sendError(w, 400)
			return
		}

		//What to send to user
		switch searchType {
		case "artist":
			sendArtist(w, r, toSearch)
		case "member":
			sendMember(w, r, toSearch)
		case "location":
			sendLocation(w, r, toSearch)
		case "firstAlbum":
			sendFirstAlbum(w, r, toSearch)
		case "creationDate":
			sendCreationDate(w, r, toSearch)
		}
	}
}
