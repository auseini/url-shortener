package main 

import (
    "log"
    "net/http"
    "github.com/auseini/url-shortener/server/routes"
)



func main(){
    http.HandleFunc("/", routes.HomeHandler)
    http.HandleFunc("/shorten",routes.ShortenHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
