package model

// Cryptoresponse is exported, it models the data we receive.
type AirportItems []struct {
   UID              string    `json:"uid"`
   NAME             string    `json:"name"`
   IATA             string    `json:"iata"`
   ICAO             string    `json:"icao"`
   LAT              float32   `json:"lat"`
   LNG              float32   `json:"lng"`
   ALT              int       `json:"alt"`
}