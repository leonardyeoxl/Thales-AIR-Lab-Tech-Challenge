package main

import (
    "fmt"
    "log"

    "github.com/leonardyeoxl/Thales-AIR-Lab-Tech-Challenge/backend/client"
)

func main() {
    airports, err := client.FetchAirports()
    if err != nil {
        log.Println(err)
      }

    fmt.Println(airports)
}