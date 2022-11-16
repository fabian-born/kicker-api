package main

import (
	"database/sql"
	b64 "encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	// "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)

func getKicker(c *gin.Context) {
	var kickerlist []Kicker
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

		rows.Scan(&id, &name)
		kickerlist = append(kickerlist, Kicker{id, name})
	}
	// c.IndentedJSON(http.StatusOK, string(queryjson))
	queryjson, _ := json.Marshal(&kickerlist)
	c.Data(http.StatusOK, "application/json", queryjson)
}

func getKickerDetail(c *gin.Context) {
	name := c.Param("name")

	var kickerlist []Kicker
	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM kicker where id = ?", name)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string

		rows.Scan(&id, &name)
		kickerlist = append(kickerlist, Kicker{id, name})
	}
	// c.IndentedJSON(http.StatusOK, string(queryjson))
	queryjson, _ := json.Marshal(&kickerlist)
	c.Data(http.StatusOK, "application/json", queryjson)
	//    c.String(http.StatusOK, "Hello %s", name)
}

func newKicker(c *gin.Context) {
	var newKicker Kicker
	//
	// curl -H "Content-Type: application/json" -X POST -d '{"ID":3,"Name":"fabian-dev2"}' http://192.168.69.22:8083/api/kicker/new
	// Call BindJSON to bind the received JSON to
	// newKicker.
	if err := c.BindJSON(&newKicker); err != nil {
		return
	}
	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	db, err := sql.Open("mysql", strings.TrimSuffix(string(b_dec_cred), "\n")+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )

	insert, err := db.Query("INSERT INTO kicker VALUES ( NULL, '" + newKicker.Name + "' )")
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	// Add the new album to the slice.
}
