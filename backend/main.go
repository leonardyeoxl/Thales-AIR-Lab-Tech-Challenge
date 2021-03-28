package main

import (
    "encoding/json"
    "fmt"

    "github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/client"
)

func main() {
    airports err:= client.FetchAirports()
    if err != nil {
        panic(err)
    }

    newsJson, err := json.Marshal(airports)
    if err != nil {
        panic(err)
    }

    fmt.Println(newsJson[0].UID)
}