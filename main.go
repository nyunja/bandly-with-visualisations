package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupie/functions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [PORT]")
		return
	}

	port := ":" + os.Args[1]
	functions.LoadData()
	http.HandleFunc("/", functions.Index)
	http.HandleFunc("/artists", functions.Artists)
	http.HandleFunc("/about", functions.About)
	http.HandleFunc("/artists/", functions.ArtistDetail)
	http.HandleFunc("/concerts", functions.Concerts)


	staticDir := "./static/"
	staticURL := "/static/"
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle(staticURL, http.StripPrefix(staticURL, fileServer))

	log.Printf("Server started at http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
