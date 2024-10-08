package functions

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
