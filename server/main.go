package main 

import (
    "fmt"
    "log"
    "net/http"
    "urlshortener/server/routes"
)

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "hellow, %s", r.URL.Path[1:])
}

func main(){
    http.HandleFunc("/", routes.ShortenHandler)
    http.HandleFunc("/redirect", routes.RedirectHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
