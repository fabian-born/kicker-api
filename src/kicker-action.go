package main

import (
	// "net/http"
        "database/sql"
        b64 "encoding/base64"
        "strings"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func KickerPlayGame(c *gin.Context) {
	id := c.Param("id")
	action := c.Param("action")

        var PlayGame Games
        if err := c.BindJSON(&PlayGame); err != nil {
              return
        }
	if  PlayGame.KickerId ==  id {
		b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	        c.JSON(200, gin.H{"status": PlayGame})
		
		if action == "startgame" {
			sqlquery, err := db.Query("insert into games value (NULL, '"+ PlayGame.KickerId +"',now(),NULL)" )
			if err != nil {panic(err.Error())}
			defer sqlquery.Close()db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
		}else if action == "endgame"{
			sqlquery, err := db.Query("UPDATE games set enddate = now() where (kickerid = '"+ PlayGame.KickerId +"' and gameid = '"+ PlayGame.GameId +"')" )
			if err != nil {panic(err.Error())}
			defer sqlquery.Close()
		}else{
			c.JSON(200, gin.H{"status": "action not found"})
		}


	        if err != nil {panic(err.Error())}
	}else{
		c.JSON(200, gin.H{"status": "Kicker not valid"})
	}
}

func KickerGoal(c *gin.Context) {
        var newGameData Gamedata
        //
	// curl -H "Content-Type: application/json" -X POST -d '{"gdid":1,"gameid":"3","kickerid":"6","goaldate":"-","teamagoal":"1","teambgoal":"0","humidity":"20","temperature":"20","teamaid":"1","teambid":"1"}' http://192.168.69.22:8084/api/kicker/goal
        if err := c.BindJSON(&newGameData); err != nil {
                return
        }
	c.JSON(200, gin.H{"status": newGameData.Gdid})
        b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
        db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )


        insert, err := db.Query("INSERT INTO gamedata  VALUES (NULL, '" + newGameData.Gameid + "', '" + string(newGameData.KickerId) + "', now(), '" + string(newGameData.TeamAGoal) + "', '" + string(newGameData.TeamBGoal) + "', '" + string(newGameData.Humidity) + "', '" + string(newGameData.Temperature) + "', '" + string(newGameData.TeamAId) + "', '" + string(newGameData.TeamBId) + "')")

        if err != nil {
                panic(err.Error())
        }

        defer insert.Close()
}

