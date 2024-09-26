package functions

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock data to be returned by the mock server
var mockArtists = []Artist{
	{
		ID:           1,
		Image:        "https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
		Name:         "Queen" ,
		Members:      []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"},
		CreationDate: 1970,
		FirstAlbum:   "14-12-1973",
		Locations:    "https://groupietrackers.herokuapp.com/api/locations/1",
		ConcertDates: "https://groupietrackers.herokuapp.com/api/dates/1",
		Relations:    "https://groupietrackers.herokuapp.com/api/relation/1",
	},
}

// Helper function to create mock HTTP server
func mockServer(response interface{}, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
	})
	return httptest.NewServer(handler)
}

// Test fetchData function
func TestFetchData_Success(t *testing.T) {
	// Create a mock server that returns mockArtists
	server := mockServer(mockArtists, http.StatusOK)
	defer server.Close()

	// Call the function we want to test, pointing it to our mock server URL
	var artists []Artist
	err := fetchData(server.URL, &artists)
	// Validate that no error occurred and that the data was fetched correctly
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	// Check if the returned data matches our mock data
	if len(artists) != 1 || artists[0].Name != mockArtists[0].Name {
		t.Errorf("Expected artist name %s, but got %s", mockArtists[0].Name, artists[0].Name)
	}
}

// Test fetchData when the server returns an error
func TestFetchData_Failure(t *testing.T) {
	// Create a mock server that returns a 500 status code
	server := mockServer(nil, http.StatusInternalServerError)
	defer server.Close()

	// Call the function and expect an error
	var artists []Artist
	err := fetchData(server.URL, &artists)

	// Ensure that an error occurred
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}

// Test LoadData function (high-level test)
func TestLoadData(t *testing.T) {
	// Create mock servers for each API endpoint
	artistServer := mockServer(mockArtists, http.StatusOK)
	defer artistServer.Close()

	// Temporarily override apiURL for testing
	originalURL := apiURL
	apiURL = artistServer.URL
	defer func() { apiURL = originalURL }()

	// Call LoadData and verify that no error occurs
	LoadData()

	// // Validate that artists were loaded successfully
	// if len(artists) != 1 || artists[0].Name != mockArtists[0].Name {
	// 	t.Errorf("Expected artist name %s, but got %s", mockArtists[0].Name, artists[0].Name)
	// }
}
