package main

import (
    "log"

    elastic "github.com/olivere/elastic/v7"
)

func main() {
    _, err := elastic.NewClient()
    if err != nil {
        log.Fatalf("Connect failed: %v", err)
    }
    log.Print("Connected")
}