package functions

import "time"

// import (
// 	"errors"
// 	"fmt"
// 	"strings"
// 	"time"
// )

type PlaceInfo struct {
	State   string
	Country string
	Date    time.Time
}

// func transformRelation(relation *Relation) error {
// 	for location, dates := range relation.DateLocs {
// 		// Split the location string into state and country
// 		splitLocation := strings.Split(location, "-")
// 		if len(splitLocation) != 2 {
// 			return errors.New("Invalid location format")
// 		}

// 		state := splitLocation[0]
// 		country := splitLocation[1]

// 		// Create a PlaceInfo struct for each date
// 		for _, dateStr := range dates {
// 			date, err := time.Parse("2006-01-02", dateStr)
// 			if err != nil {
// 				return errors.New("Error parsing date" + dateStr)
// 			}

// 			placeInfo := PlaceInfo{
// 				State:   state,
// 				Country: country,
// 				Date:    date,
// 			}
// 			fmt.Println(placeInfo)

// 			relation.Places = append(relation.Places, placeInfo)
// 		}
// 	}

// 	return nil
// }
