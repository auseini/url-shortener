package routes 

import (
    "fmt"
    "net/http"
)

func ShortenHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "comingin from routes, %s", r.URL.Path[1:])
}

func RedirectHandler(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, "https://www.google.com", http.StatusMovedPermanently);
}
