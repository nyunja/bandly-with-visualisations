package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define structs for API response
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Index []Location `json:"index"`
}
type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"locations"`
	Dates    string   `json:"dates"`
}

type Dates struct {
	Index []Date
}

type Date struct {
	ID   int      `json:"id"`
	Date []string `json:"dates"`
}

type Relations struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID       int                 `json:"id"`
	DateLocs map[string][]string `json:"datesLocations"`
}

type Data struct {
	Artists   []Artist
	Locations Locations
	Dates     Dates
	Relations Relations
}

type BandDetails struct {
	Artist   Artist
	Location Location
	Dates Date
	Relation Relation
}

var (
	artists   []Artist
	locations Locations
	dates     Dates
	relations Relations
)

var apiURL = "https://groupietrackers.herokuapp.com/api/"

func fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error fetching JSON from %s", url)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data from %s: %s", url, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: Reading from repsonse body")
		return err
	}
	return json.Unmarshal(body, target)
}

func LoadData() {
	var err error
	if err = fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	if err = fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations); err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}
	if err = fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates); err != nil {
		log.Fatalf("Error fetching dates: %v", err)
	}
	if err = fetchData("https://groupietrackers.herokuapp.com/api/relation", &relations); err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}
}
