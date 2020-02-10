package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type APIstruct struct {
	ID        int
	Artists   Artists
	Locations Locations
	Dates     Dates
	Relation  Relation
}

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type Relation struct {
	Index []struct {
		ID             int64               `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

var API APIstruct

func main() {
	parseJSON()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", indexHandle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening server at port %v\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ParseArtists(url string) (Artists, error) {
	var data Artists
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	return data, err
}

func ParseLocations(url string) (Locations, error) {
	var data Locations
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	return data, err
}

func ParseDates(url string) (Dates, error) {
	var data Dates
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	return data, err
}

func ParseRelation(url string) (Relation, error) {
	var data Relation
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	json.Unmarshal(body, &data)
	return data, err
}

func parseJSON() {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocations := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"
	API.Artists, _ = ParseArtists(urlArtists)
	API.Locations, _ = ParseLocations(urlLocations)
	API.Dates, _ = ParseDates(urlDates)
	API.Relation, _ = ParseRelation(urlRelation)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		temp, err := template.ParseFiles("./static/templates/index.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, API)
	case "POST":
		toSearch := r.FormValue("toSearch")
		searchType := r.FormValue("searchType")

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

func sendArtist(w http.ResponseWriter, r *http.Request, toSearch string) {
	API.ID = -1
	for i := 0; i < 52; i++ {
		if strings.ToLower(API.Artists[i].Name) == strings.ToLower(toSearch) {
			API.ID = i
			break
		}
	}
	if API.ID == -1 {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		temp, err := template.ParseFiles("./static/templates/artist.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, API)
	}
}

func sendMember(w http.ResponseWriter, r *http.Request, toSearch string) {
	type MemberPage struct {
		Title  string
		Artist []string
	}
	var memberPage MemberPage
	for i := 0; i < 52; i++ {
		for _, member := range API.Artists[i].Members {
			if strings.ToLower(member) == strings.ToLower(toSearch) {
				memberPage.Title = member + "<br>is a member of"
				memberPage.Artist = append(memberPage.Artist, API.Artists[i].Name)
			}
		}
	}
	if memberPage.Title == "" {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		temp, err := template.ParseFiles("./static/templates/member.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, memberPage)
	}
}

func sendLocation(w http.ResponseWriter, r *http.Request, toSearch string) {
	type LocationPage struct {
		Title   string
		Artists []string
	}
	var locationPage LocationPage

	for i, all := range API.Locations.Index {
		for _, location := range all.Locations {
			if strings.ToLower(location) == strings.ToLower(toSearch) {
				locationPage.Title = "Concerts in " + location
				locationPage.Artists = append(locationPage.Artists, API.Artists[i].Name)
				break
			}
		}
	}
	if locationPage.Title == "" {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		temp, err := template.ParseFiles("./static/templates/location.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, locationPage)
	}
}

func sendFirstAlbum(w http.ResponseWriter, r *http.Request, toSearch string) {
	type FirstAlbumPage struct {
		Title   string
		Artists []string
	}
	var firstAlbumPage FirstAlbumPage

	for i, artist := range API.Artists {
		if strings.ToLower(artist.FirstAlbum) == strings.ToLower(toSearch) {
			firstAlbumPage.Title = "Artists / Bands relased their first album in " + artist.FirstAlbum
			firstAlbumPage.Artists = append(firstAlbumPage.Artists, API.Artists[i].Name)
		}
	}
	if firstAlbumPage.Title == "" {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		temp, err := template.ParseFiles("./static/templates/firstalbum.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, firstAlbumPage)
	}
}

func sendCreationDate(w http.ResponseWriter, r *http.Request, toSearch string) {
	year, _ := strconv.Atoi(toSearch)
	type CreationDatePage struct {
		Title   string
		Artists []string
	}
	var creationDatePage CreationDatePage

	for i, artist := range API.Artists {
		if artist.CreationDate == year {
			creationDatePage.Title = "Artists / Bands created in " + toSearch
			creationDatePage.Artists = append(creationDatePage.Artists, API.Artists[i].Name)
		}
	}
	if creationDatePage.Title == "" {
		temp, err := template.ParseFiles("./static/templates/noresult.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, toSearch)
	} else {
		temp, err := template.ParseFiles("./static/templates/creationdate.html")
		if err != nil {
			http.Error(w, "500 internal server error.", http.StatusInternalServerError)
			return
		}
		temp.Execute(w, creationDatePage)
	}
}
