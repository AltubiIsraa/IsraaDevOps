package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
)

type Simple struct {
        Name string
        Description string
        Url string
}

func SimpleFactory(host string) simple1 {
	return Simple{"Hello","Israa", host}
}

func handler(w http.ResponseWriter, r *http.Request) {
        simple := SimpleFactory(r.Host)

        jsonOutput, _ := json.Marshal(simple1)

        fmt.Fprintln(w, string(jsonOutput))
}

func main() {
        fmt.Println("Server started")
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":9070", nil))
}
