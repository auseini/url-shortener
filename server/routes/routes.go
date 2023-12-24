package routes 

import (
    "strings"
    "fmt"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request){
    shortId := strings.TrimPrefix(r.URL.Path, "/")

    if shortId != "" {
        redirectHandler(w, r, shortId)
    } else{
        fmt.Fprintf(w, "will render some html here")
    }


}

func ShortenHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "comingin from routes, %s", r.URL.Path[1:])
}

func redirectHandler(w http.ResponseWriter, r *http.Request, shortId string){
    http.Redirect(w, r, "https://www.google.com", http.StatusMovedPermanently);
}
