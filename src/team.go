package main

import (
        "database/sql"
        b64 "encoding/base64"
        "encoding/json"
        "log"
        "net/http"
        "strings"

        "github.com/gin-gonic/gin"
        _ "github.com/go-sql-driver/mysql"
)


// team functions //
func ListTeams(c *gin.Context) {
	var teamlist []Team

	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
	defer db.Close()
	rows, err := db.Query("SELECT id, name, teamtype FROM teams")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
                var teamtype string

		rows.Scan(&id, &name)
		teamlist = append(teamlist, Team{id, name, teamtype})
	}
	queryjson, _ := json.Marshal(&teamlist)
	c.Data(http.StatusOK, "application/json", queryjson)
}

func ListAllPlayer(c *gin.Context) {
	var playerlist []Player

	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
	defer db.Close()
	rows, err := db.Query("SELECT id, surename, lastname FROM player")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var surename string
        var lastname string

		rows.Scan(&id, &surename, &lasname)
		playerlist = append(playerlist, Player{id, surename, lastname})
	}
	queryjson, _ := json.Marshal(&playerlist)
	c.Data(http.StatusOK, "application/json", queryjson)
}

func AddTeam(c *gin.Context) {
    var newTeam Team

	if err := c.BindJSON(&newTeam); err != nil {
		return
	}
	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
	insert, err := db.Query("INSERT INTO teams VALUES ( NULL, '" + newTeam.Name + newTeam.Teamtype + "' )")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
 }
func DelTeam(c *gin.Context) { }
func UpdTeam(c *gin.Context) { }
func AddTeamMember(c *gin.Context) { }