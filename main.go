package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//Parsing JSON
	parseJSON()

	//Serving static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//Index handler
	http.HandleFunc("/", indexHandle)

	//Set port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Starting server
	fmt.Printf("Listening server at port %v\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
