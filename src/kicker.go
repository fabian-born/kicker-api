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


func getKicker(c *gin.Context) {
    type Kicker struct {
       ID   int    `json:"id"`
       Name string `json:"name"`
    }
    var kickerlist []Kicker


    b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
    db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred),"\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/smartkicker")

    defer db.Close()
    rows, err := db.Query("SELECT id, name FROM kicker")
    if err != nil {
      log.Fatal(err)
    }
    for rows.Next() {
       var id int
       var name string

       rows.Scan(&id ,&name)
       kickerlist = append(kickerlist, Kicker{id, name })
    }
    // c.IndentedJSON(http.StatusOK, string(queryjson))
    queryjson, _ := json.Marshal(&kickerlist)
    c.Data(http.StatusOK, "application/json", queryjson)
}

func getKickerDetail(c *gin.Context) {
    id := c.Param("id")
    type Kicker struct {
       ID   int    `json:"id"`
       Name string `json:"name"`
    }
    var kickerlist []Kicker
    b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
    db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred),"\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/smartkicker")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query("SELECT id, name FROM kicker where id = ?",id)
    if err != nil {
      log.Fatal(err)
    }
    for rows.Next() {
       var id int
       var name string

       rows.Scan(&id ,&name)
       kickerlist = append(kickerlist, Kicker{id, name })
    }
    // c.IndentedJSON(http.StatusOK, string(queryjson))
    queryjson, _ := json.Marshal(&kickerlist)
    c.Data(http.StatusOK, "application/json", queryjson)
    //    c.String(http.StatusOK, "Hello %s", name)
}

func actionKicker(c *gin.Context) {
  id := c.Param("id")
  action := c.Param("action")
  message := id + " is " + action
  c.String(http.StatusOK, message)
}

