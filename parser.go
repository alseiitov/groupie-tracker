package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Main struct
var All API

//Parse JSON from URL and write to main struct
func parseJSON() {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocations := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"
	All.Artists, _ = ParseArtists(urlArtists)
	All.Locations, _ = ParseLocations(urlLocations)
	All.Dates, _ = ParseDates(urlDates)
	All.Relation, _ = ParseRelation(urlRelation)
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
