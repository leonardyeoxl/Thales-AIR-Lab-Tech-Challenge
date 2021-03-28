package main

import (
    "encoding/json"
    "fmt"

    "github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/client"
)

func main() {
    airports := client.FetchAirports()
    newsJson, err := json.Marshal(airports)
    if err != nil {
        panic(err)
    }

    fmt.Println(newsJson)
}