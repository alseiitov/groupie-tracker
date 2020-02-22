package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
