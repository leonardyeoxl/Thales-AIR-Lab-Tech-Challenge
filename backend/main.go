package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/client"
)

func main() {
    airports, err := client.FetchAirports()
    newsJson, err := json.Marshal(airports)
    if err != nil {
        panic(err)
    }

    fmt.Println(newsJson)
}