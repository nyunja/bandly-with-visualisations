package functions

import (
	"encoding/json"
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
	if id < 1 || id > 52 {
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

// func Search(w http.ResponseWriter, r *http.Request) {
// 	query := strings.ToLower(r.URL.Query().Get("q"))

// 	var results []Artist

// 	for _, artist := range artists {
// 		if strings.Contains(strings.ToLower(artist.Name), query) {
// 			results = append(results, artist)
// 		}
// 	}

// 	data := map[string]interface{}{
// 		"Artists": results,
// 	}

// 	tmpl, err := template.ParseFiles("templates/search.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

func Search(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))

	var results []SearchResult

	for j, artist := range artists {
		lowerCaseMembers := make([]string, len(artist.Members))
		for j, member := range artist.Members {
			lowerCaseMembers[j] = strings.ToLower(member)
		}
		result := SearchResult{}
		if strings.Contains(strings.ToLower(artist.Name), query) {
			result.ID = artist.ID
			result.Match = artist.Name + " - artist/band"
			if result.ID > 0 {
				results = append(results, result)
			}
			result = SearchResult{}
		}
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), query) {
			result.ID = artist.ID
			result.Match = strconv.Itoa(artist.CreationDate) + " - " + artist.Name + "(Creation Date)"
			if result.ID > 0 {
				results = append(results, result)
			}
			result = SearchResult{}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			result.ID = artist.ID
			result.Match = artist.FirstAlbum + " - " + artist.Name + "(First Album)"
			if result.ID > 0 {
				results = append(results, result)
			}
			result = SearchResult{}
		}
		for i, member := range lowerCaseMembers {
			if strings.Contains(member, query) {
				result.ID = artist.ID
				result.Match = artist.Members[i] + " - member"
				if result.ID > 0 {
					results = append(results, result)
				}
				result = SearchResult{}
			}
		}
		for _, location := range locations.Index[j].Location {
			if strings.Contains(location, query) {
				result.ID = artist.ID
				result.Match = location+ " - " + artist.Name + "(Location)"
				if result.ID > 0 {
					results = append(results, result)
				}
				result = SearchResult{}
			}
		}
		// 	strings.Contains(strings.Join(lowerCaseMembers, " "), query) ||
		// 	strings.Contains(strings.Join(locations.Index[i].Location, " "), query) ||
		// 	strings.Contains(strings.Join(dates.Index[i].Date, " "), query) ||
		// 	strings.Contains(fmt.Sprintf("%v", relations.Index[i].DateLocs), query) {

		// 	result := SearchResult{
		// 		ID:        artist.ID,
		// 		Name:      artist.Name,
		// 		Members:   artist.Members,
		// 		FirstAlbum: artist.FirstAlbum,
		// 		CreationDate: artist.CreationDate,
		// 		Location:  strings.Join(locations.Index[i].Location, ", "),
		// 		Dates:     strings.Join(dates.Index[i].Date, ", "),
		// 		Relations: fmt.Sprintf("%v", relations.Index[i].DateLocs),
		// 	}
		// 	results = append(results, result)
		// }
		

	}

	// Return the search results as JSON
	json.NewEncoder(w).Encode(results)
}
