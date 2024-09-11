package functions

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Response struct {
	pageTitle string
	Data      Data
}

var dataRes Response

type BandResponse struct {
	pageTitle string
	Band      BandDetails
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Println("OK: ", http.StatusOK)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataRes = Response{
		pageTitle: "Artists",
		Data: Data{
			Artists:   artists,
			Locations: locations,
			Dates:     dates,
			Relations: relations,
		},
	}
	tmpl.Execute(w, dataRes)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if strings.ToUpper(r.Method) != http.MethodGet {
		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataRes = Response{
		pageTitle: "Artists",
		Data: Data{
			Artists:   artists,
			Locations: locations,
			Dates:     dates,
			Relations: relations,
		},
	}
	tmpl.Execute(w, dataRes)
}

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/about.html")
		if err != nil {
			ServeError(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else {
		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func Concerts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/concerts" {
		ServeError(w, "Page not found", http.StatusNotFound)
	}

	if strings.ToUpper(r.Method) != http.MethodGet {
		ServeError(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	tmpl, err := template.ParseFiles("templates/concerts.html")
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	dataConcerts := Response{
		pageTitle: "Concerts",
		Data: Data{
			Artists:   artists,
			Locations: locations,
			Dates:     dates,
			Relations: relations,
		},
	}

	tmpl.Execute(w, dataConcerts)
}

func ArtistDetail(w http.ResponseWriter, r *http.Request) {
	// Extract the artist ID from the URL (e.g., /artists/{id})
	idStr := strings.TrimPrefix(r.URL.Path, "/artists/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ServeError(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	index := id - 1
	if id == 0 {
		ServeError(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Load the artist detail template
	tmpl, err := template.ParseFiles("templates/band.html")
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Prepare the response data
	data := BandResponse{
		pageTitle: artists[index].Name,
		Band: BandDetails{
			Artist:   artists[index],
			Location: locations.Index[index],
			Dates:    dates.Index[index],
			Relation: relations.Index[index],
		},
	}

	tmpl.Execute(w, data)
}
