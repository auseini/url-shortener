package routes 

import (
    "encoding/json"
    "strings"
    "fmt"
    "net/http"
)


type ShortenRequest struct {
    Url string
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
    shortId := strings.TrimPrefix(r.URL.Path, "/")

    if shortId != "" {
        redirectHandler(w, r, shortId)
    } else{
        fmt.Fprintf(w, "will render some html here")
    }


}

func ShortenHandler(w http.ResponseWriter, r *http.Request){

    method := r.Method

    if method != "POST" {
        fmt.Fprintf(w, "only POST allowed")
        return
    }

    var request ShortenRequest

    err := json.NewDecoder(r.Body).Decode(&request) 

    if err != nil {
        fmt.Fprintln(w, "Could not read body")
        return
    }

    url := request.Url
    fmt.Fprintf(w, "%s", url) 
}

func redirectHandler(w http.ResponseWriter, r *http.Request, shortId string){
    http.Redirect(w, r, "https://www.google.com", http.StatusMovedPermanently);
}
