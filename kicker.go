package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
)


func getKicker(c *gin.Context) {
    type Kicker struct {
       ID   int    `json:"id"`
       Name string `json:"name"`
    }
    var kickerlist []Kicker


    db, err := sql.Open("mysql", "root:root@tcp(192.168.69.22:8989)/smartkicker")
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
    db, err := sql.Open("mysql", "root:root@tcp(192.168.69.22:8989)/smartkicker")
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

