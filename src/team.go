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
	rows, err := db.Query("SELECT id, name FROM kicker")
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
	// c.IndentedJSON(http.StatusOK, string(queryjson))
	queryjson, _ := json.Marshal(&teamlist)
	c.Data(http.StatusOK, "application/json", queryjson)
}

func AddTeam(c *gin.Context) {
        var newTeam Team
	//
	// curl -H "Content-Type: application/json" -X POST -d '{"ID":1,"Name":"fabian-dev2"}' http://192.168.69.22:8083/api/kicker/new
	// Call BindJSON to bind the received JSON to
	// newKicker.
	if err := c.BindJSON(&newTeam); err != nil {
		return
	}
	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )

	insert, err := db.Query("INSERT INTO teams VALUES ( NULL, '" + newTeam.Name + "' )")
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	// Add the new album to the slice.
 }
func DelTeam(c *gin.Context) { }
func UpdTeam(c *gin.Context) { }
