package model

import (
   "fmt"
)

// Cryptoresponse is exported, it models the data we receive.
type AirportItem struct {
   UID              string    `json:"uid"`
   NAME             string    `json:"name"`
   IATA             string    `json:"iata"`
   ICAO             string    `json:"icao"`
   LAT              float32   `json:"lat"`
   LNG              float32   `json:"lng"`
   ALT              int       `json:"alt"`
}

type Airport []AirportItem

//TextOutput is exported,it formats the data to plain text.
func (airport AirportItem) TextOutput() string {
   response_out := fmt.Sprintf(
   "UID: %s\nNAME : %s\nIATA: %s\nICAO: %s\nLAT: %f\nLNG: %f\nALT: %d\n",
   airport[0].UID, airport[0].NAME, airport[0].IATA, airport[0].ICAO, airport[0].LAT, airport[0].LNG, airport[0].ALT)
   return response_out
}