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
    db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred),"\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
    // return games data
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

func getKickerlatestGame(c *gin.Context) {
    id := c.Param("name") 
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
    db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred),"\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
    // return games data
    defer db.Close()
    rows, err := db.Query("SELECT * FROM games where kickerid = " + id + " ORDER BY gameid DESC LIMIT 1")
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

func getGameData(c *gin.Context) {
        name := c.Param("id")

        var gamedatalist []Gamedata
        b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
        db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
        if err != nil {
                panic(err.Error())
        }
        defer db.Close()

        rows, err := db.Query("SELECT * FROM gamedata where gamesid = ?", name)
        if err != nil {
                log.Fatal(err)
        }
        for rows.Next() {
                var gdid int 
		var gameid string
		var kickerid string
                var goaldate string
		var teamagoal string
		var teambgoal string
		var humidity string
		var temperature string
		var teamaid string
		var teambid string

                rows.Scan(&gdid,&gameid,&kickerid,&goaldate,&teamagoal,&teambgoal,&humidity,&temperature,&teamaid,&teambid)
		gamedatalist = append(gamedatalist, Gamedata{gdid,gameid,kickerid,goaldate,teamagoal,teambgoal,humidity,temperature,teamaid,teambid})
        }
        // c.IndentedJSON(http.StatusOK, string(queryjson))
        queryjson, _ := json.Marshal(&gamedatalist)
        c.Data(http.StatusOK, "application/json", queryjson)
        //    c.String(http.StatusOK, "Hello %s", name)
}

