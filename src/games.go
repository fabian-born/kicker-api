package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
     b64 "encoding/base64"
     "strings"
)


func getGames(c *gin.Context) {
    type Kicker struct {
       ID   int    `json:"id"`
       Name string `json:"name"`
    }
    type Games struct {
       GID  int    `json:"gameid"`
       KID  int `json:"kickerid"`
       startdate string `json:"startdate"`
       enddate string `json:"enddate"`
    }
    var gamelist []Games

    b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
    db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred),"\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/smartkicker")

    defer db.Close()
    rows, err := db.Query("SELECT * FROM games")
    if err != nil {
      log.Fatal(err)
    }
    for rows.Next() {
       var gameid int
       var kickerid int
       var startdate string
       var enddate string

       rows.Scan(&gameid ,&kickerid, &startdate, &enddate)
       gamelist = append(gamelist, Games{gameid, kickerid, startdate, enddate })
    }
    queryjson, _ := json.Marshal(&gamelist)
    c.Data(http.StatusOK, "application/json", queryjson)
}

