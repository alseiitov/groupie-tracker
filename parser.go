package main

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
