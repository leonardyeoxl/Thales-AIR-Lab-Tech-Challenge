package main

import (
    "encoding/json"
    "log"
    "fmt"

    "github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/client"
)

func main() {
    airports, err:= client.FetchAirports()
    if err != nil {
        log.Fatalln(err)
    }

    airports_json, err := json.Marshal(airports)
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(string(airports_json))
}