package client

import (
   "encoding/json"
   "fmt"
   "log"
   "net/http"

   "github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/model"
)

func FetchAirports() (string, error) {
	//Build The URL string	
   	URL := "https://open-atms.airlab.aero/api/v1/airac/airports"
	//We make HTTP request using the Get function
	req, err := http.NewRequest("GET", URL, nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("Content-Type", contentType)
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

	//Create a variable of the same type as our model
   var airportResp model.AirportResponse
	//Decode the data
   if err := json.NewDecoder(resp.Body).Decode(&airportResp); err != nil {
      log.Fatal(err)
   }
	//Invoke the text output function & return it with nil as the error value
   return airportResp.TextOutput(), nil
}