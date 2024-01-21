package routes

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/auseini/url-shortener/server/db"
	"github.com/auseini/url-shortener/server/templates"
	"github.com/redis/go-redis/v9"
	"github.com/teris-io/shortid"
)


func HomeHandler(w http.ResponseWriter, r *http.Request){
    shortId := strings.TrimPrefix(r.URL.Path, "/")

    if shortId != "" {
        redirectHandler(w, r, shortId)
    } else{
        templates.HomePage().Render(r.Context(), w) 
    }
}

func ShortenHandler(w http.ResponseWriter, r *http.Request){

    method := r.Method

    if method != "POST" {
        fmt.Fprintf(w, "only POST allowed")
        return
    }

    url := r.FormValue("url") 
    if url == "" {
        fmt.Fprintf(w,"no link recieved")
        return
    }


    rdb := db.CreateDbClient()

    sid, err := shortid.New(1, shortid.DefaultABC, 2342)
    if err != nil {
        panic(err)
    }

    shortId, err := sid.Generate()

    if err != nil {
        panic(err)
    }

    shortId = shortId[:5]
    if url[:8] != "https://" {
        url = "https://" + url
    }
    
    err = rdb.Set(db.Ctx, shortId, url, 24 * time.Hour).Err()
    if err != nil {
        panic(err)
    }

    redirectUrl := os.Getenv("REDIRECT_URL")
    templates.Link(redirectUrl + shortId).Render(r.Context(), w)
}

func redirectHandler(w http.ResponseWriter, r *http.Request, shortId string){
    rdb := db.CreateDbClient()

    val, err := rdb.Get(db.Ctx, shortId).Result()
    if err == redis.Nil {
        fmt.Fprintf(w, "url not found")
        return
    } else if err != nil {
        panic(err)
    }

    if val == "" {
        fmt.Fprintf(w, "url not found")
        return    
    }


    http.Redirect(w, r, val, http.StatusSeeOther)

}
