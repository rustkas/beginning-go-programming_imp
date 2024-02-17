package main

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
    fmt.Println("Network Requests Demo")

    // Retry the request up to 3 times
    for attempts := 1; attempts <= 3; attempts++ {
        response, err := http.Get(url)
        if err != nil {
            fmt.Printf("Error making request: %v\n", err)
            fmt.Printf("Attempt %d\n", attempts)
            time.Sleep(time.Second * 2) // Wait before retrying
            continue
        }
        defer response.Body.Close()

        fmt.Printf("Response Type: %T\n", response)

        // Read the response body
        bytes, err := io.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("Error reading response body: %v\n", err)
            fmt.Printf("Attempt %d\n", attempts)
            time.Sleep(time.Second * 2) // Wait before retrying
            continue
        }
        content := string(bytes)
        fmt.Print(content)
        return // Success, exit the loop
    }

    fmt.Println("Failed after 3 attempts")
}
