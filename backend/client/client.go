package client

import (
   "log"
   "net/http"
   "io/ioutil"
   "fmt"
   "encoding/json"

   "github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/model"
)

func FetchAirports() ([]model.Airport, error) {
   contentType := "application/json"
	// Build The URL string	
   URL := "https://open-atms.airlab.aero/api/v1/airac/airports"
	// Make HTTP request using the Get function
	req, err := http.NewRequest("GET", URL, nil)
   if err != nil {
      log.Fatal(err)
   }
   req.Header.Set("Content-Type", contentType)
   req.Header.Set("api-key", "lgBaEkJ1TLQrwFDhtwe2mqLWIgoiyxue9kmrNkvOKpdjfhyXHIcdw7MNLmTLopH6")
   resp, err := http.DefaultClient.Do(req)
   if err != nil {
      log.Fatal(err)
   }
   defer resp.Body.Close()

   bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bytes))

   // airports := []model.Airport{}
   airports := make([]model.Airport{},0)
   // err = json.NewDecoder(resp.Body).Decode(&airports)
   err = json.Unmarshal(bytes, &airports)
   if err != nil {
		log.Fatalln(err)
	}
   fmt.Printf("%s", airports)

   return airports, nil
}