package main

import (
	"net/http"
        "database/sql"
        b64 "encoding/base64"
        "strings"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func KickerStartgame(c *gin.Context) {
	name := c.Param("name")
	message := "comming soon! " + name + " action: StartGame"
	c.String(http.StatusOK, message)

}

func KickerGoal(c *gin.Context) {
        var newGameData Gamedata
        //
	// curl -H "Content-Type: application/json" -X POST -d '{"gdid":1,"gameid":"3","kickerid":"6","goaldate":"-","teamagoal":"1","teambgoal":"0","humidity":"20","temperature":"20","teamaid":"1","teambid":"1"}' http://192.168.69.22:8084/api/kicker/goal
        // Call BindJSON to bind the received JSON to
        // newKicker.
        if err := c.BindJSON(&newGameData); err != nil {
                return
        }
	c.JSON(200, gin.H{"status": newGameData.Gdid})
        b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
        db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/smartkicker")


        insert, err := db.Query("INSERT INTO gamedata  VALUES (NULL, '" + newGameData.Gameid + "', '" + string(newGameData.KickerId) + "', now(), '" + string(newGameData.TeamAGoal) + "', '" + string(newGameData.TeamBGoal) + "', '" + string(newGameData.Humidity) + "', '" + string(newGameData.Temperature) + "', '" + string(newGameData.TeamAId) + "', '" + string(newGameData.TeamBId) + "')")

        if err != nil {
                panic(err.Error())
        }

        // be careful deferring Queries if you are using transactions
        defer insert.Close()
        // Add the new album to the slice.
}

func KickerEndgame(c *gin.Context) {
	name := c.Param("name")
	message := "comming soon! " + name + " action: StartGame"
	c.String(http.StatusOK, message)

}
