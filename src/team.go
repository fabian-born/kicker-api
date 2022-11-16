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
func ListTeams(c *gin.Context) { }

func AddTeam(c *gin.Context) {
        var newTeam Team
	//
	// curl -H "Content-Type: application/json" -X POST -d '{"ID":1,"Name":"fabian-dev2"}' http://192.168.69.22:8083/api/kicker/new
	// Call BindJSON to bind the received JSON to
	// newKicker.
	if err := c.BindJSON(&newKicker); err != nil {
		return
	}
	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.dbkicker )

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