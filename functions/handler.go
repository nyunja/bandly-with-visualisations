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

func Index(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != "/" {
		ServeError(w, "Page not found", http.StatusNotFound, "templates/error.html")
		return
	}

	if r.Method == http.MethodGet {
		fmt.Println("OK: ", http.StatusOK)
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(err)
		ServeError(w, "Internal server error", http.StatusInternalServerError, "templates/error.html")
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

// func Artists(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/artists" {
// 		ServeError(w, "Page not found", http.StatusNotFound)
// 		return
// 	}

// 	if strings.ToUpper(r.Method) != http.MethodGet {
// 		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	tmpl, err := template.ParseFiles("templates/artists.html")
// 	if err != nil {
// 		ServeError(w, "Internal server error", http.StatusInternalServerError)
// 		return
// 	}

// 	dataRes = Response{
// 		pageTitle: "Artists",
// 		Data: Data{
// 			Artists:   artists,
// 			Locations: locations,
// 			Dates:     dates,
// 			Relations: relations,
// 		},
// 	}
// 	tmpl.Execute(w, dataRes)
// }

func About(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != "/about" {
		ServeError(w, "Page not found", http.StatusNotFound, "templates/error.html")
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles(path)
		if err != nil {
			ServeError(w, "Internal server error", http.StatusInternalServerError, "templates/error.html")
			return
		}
		tmpl.Execute(w, nil)
	} else {
		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed, "templates/error.html")
	}
}

func ArtistDetail(w http.ResponseWriter, r *http.Request, path string) {
	// Extract the artist ID from the URL (e.g., /artists/{id})
	idStr := strings.TrimPrefix(r.URL.Path, "/artists/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ServeError(w, "Invalid artist ID", http.StatusBadRequest, "templates/error.html")
		return
	}

	index := id - 1
	if id < 1 || id > 52 {
		ServeError(w, "Artist not found", http.StatusNotFound, "templates/error.html")
		return
	}

	// Load the artist detail template
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError, "templates/error.html")
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

// func FetchSuggestions(query string) ([]SearchResult, error) {
// 	url := fmt.Sprintf("%s?q=%s", apiURL, query)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("error fetching suggestions: %s", resp.Status)
// 	}

// 	var data struct {
// 		Artists []Artist `json:"artists"`
// 	}

// 	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
// 		return nil, err
// 	}

// 	var suggestions []SearchResult

// 	for _, artist := range data.Artists {
// 		suggestions = append(suggestions, SearchResult{
// 			ID:        artist.ID,
// 			Name:      artist.Name,
// 			Members:   artist.Members,
// 			Location:  artist.Locations,
// 			Dates:     artist.ConcertDates,
// 			Relations: artist.Relations,
// 		})
// 	}
// 	return suggestions, nil
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
			result.Name = artist.Name
			result.Match = artist.Name + " - artist/band"
			if result.ID > 0 {
				results = append(results, result)
			}
			result = SearchResult{}
		}
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), query) {
			result.ID = artist.ID
			result.Name = artist.Name
			result.Match = strconv.Itoa(artist.CreationDate) + " - " + artist.Name + "(Creation Date)"
			if result.ID > 0 {
				results = append(results, result)
			}
			result = SearchResult{}
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			result.ID = artist.ID
			result.Name = artist.Name
			result.Match = artist.FirstAlbum + " - " + artist.Name + "(First Album)"
			if result.ID > 0 {
				results = append(results, result)
			}
			result = SearchResult{}
		}
		for i, member := range lowerCaseMembers {
			if strings.Contains(member, query) {
				result.ID = artist.ID
				result.Name = artist.Name
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
				result.Name = artist.Name
				result.Match = location + " - " + artist.Name + "(Location)"
				if result.ID > 0 {
					results = append(results, result)
				}
				result = SearchResult{}
			}
		}
	}

	// Return the search results as JSON
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
