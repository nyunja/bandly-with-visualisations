package functions

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchData(t *testing.T) {
	var artists []Artist
	fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if artists[0].Name != "Queen" {
		t.Errorf("fetch function returned wrong artists: expected %v got %v", "Queen", artists[0].Name)
	}
}
func TestServeError(t *testing.T) {
	// Create a response to our handler
	w := httptest.NewRecorder()
	ServeError(w, "Internal Server Error", http.StatusInternalServerError, "../templates/error.html")
	// Check the status code
	if status := w.Code; status != http.StatusInternalServerError {
        t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusInternalServerError)
    }
}
func TestIndex(t *testing.T) {
	testData()
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	resRecorder := httptest.NewRecorder()

	// Call the handler function directly
	Index(resRecorder, req, "../templates/index.html")

	// Check the status code
	if status := resRecorder.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func TestAbout(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	resRecorder := httptest.NewRecorder()

	// Call the handler function directly
	About(resRecorder, req, "../templates/about.html")

	// Check the status code
	if status := resRecorder.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func TestArtistDetail(t *testing.T) {
	testData()
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/artists/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	resRecorder := httptest.NewRecorder()

	// Call the handler function directly
	ArtistDetail(resRecorder, req, "../templates/band.html")

	// Check the status code
	if status := resRecorder.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func TestSearch(t *testing.T) {
	testData()
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/search?=h", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	resRecorder := httptest.NewRecorder()

	// Call the handler function directly
	Search(resRecorder, req)

	// Check the status code
	if status := resRecorder.Code; status != http.StatusOK{
		t.Errorf("handler returned wrong status code: got %v expected %v", status, http.StatusOK)
	}
}

func testData() {
	// Mock data
	artists = []Artist{
		{ID: 1, Name: "John Paul"},
	}
	locations = Locations{
		Index: []Location{
			{ID: 1, Location: []string{"Kisumu"}, Dates: "2024"},
		},
	}
	dates = Dates{
		Index: []Date{
			{ID: 1, Date: []string{"2024"}},
		},
	}
	relations = Relations{
		Index: []Relation{
			{ID: 1, DateLocs: map[string][]string{
				"Kisumu": {"2024"},
			}},
		},
	}
}
