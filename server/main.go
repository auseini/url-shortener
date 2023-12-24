package main 

import (
    "log"
    "net/http"
    "urlshortener/server/routes"
)

func main(){
    http.HandleFunc("/", routes.HomeHandler)
    http.HandleFunc("/shorten",routes.ShortenHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
