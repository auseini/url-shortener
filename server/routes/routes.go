package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
    "github.com/teris-io/shortid"
	"github.com/auseini/url-shortener/server/db"
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


    rdb := db.CreateDbClient()

    sid, err := shortid.New(1, shortid.DefaultABC, 2342)
    if err != nil {
        panic(err)
    }

    shortId, err := sid.Generate()

    if err != nil {
        panic(err)
    }

    shortId = shortId[:4]
    url = "https://" + url
    
    err = rdb.Set(db.Ctx, shortId, url, 24 * time.Hour).Err()
    if err != nil {
        panic(err)
    }

    fmt.Fprintf(w, "%s", shortId) 
}

func redirectHandler(w http.ResponseWriter, r *http.Request, shortId string){
    rdb := db.CreateDbClient()

    val, err := rdb.Get(db.Ctx, shortId).Result()
    if err != nil {
        panic(err)
    }

    if val == "" {
        fmt.Fprintf(w, "url not found")
        return    
    }


    http.Redirect(w, r, val, http.StatusSeeOther)

}
