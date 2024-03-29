package main

import (
	"database/sql"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-yaml/yaml"
)

type Kicker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Games struct {
	GameId    string    `json:"gameid"`
	KickerId  string    `json:"kickerid"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}

type Gamedata struct {
	Gdid        int     `json:"gdid"`
	Gameid      string     `json:"gameid"`
	KickerId    string    `json:"kickerid"`
	Goaldate    string  `json:"goaldate"`
	TeamAGoal   string     `json:"teamagoal"`
	TeamBGoal   string     `json:"teambgoal"`
	Humidity    string `json:"humidity"`
	Temperature string `json:"temperature"`
	TeamAId     string     `json:"teamaid"`
	TeamBId     string     `json:"teambid"`
}
type conf struct {
	DBHost     string `yaml:"dbhost"`
	DBPort     string `yaml:"dbport"`
	Credential string `yaml:"dbcredential"`
	Dbkicker   string `yaml:"database"`
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Teamtype string `json:"teamtype"` // single, multi
}

type Player struct {
	ID   int    `json:"id"`
	Surename string `json:"surename"`
	Lastname string `json:"lastname"`
}

func (c *conf) GetConfig() *conf {
	confContent, err := ioutil.ReadFile("/app/config/config.yml")
	// confContent, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	confContent = []byte(os.ExpandEnv(string(confContent)))
	// c := &Config{}
	if err := yaml.Unmarshal(confContent, c); err != nil {
		panic(err)
	}
	return c
}



var myconf conf

func main() {
	//    var myconf conf
	myconf.GetConfig()
	// expand environment variables
	b_dec_cred, _ := b64.StdEncoding.DecodeString((myconf.Credential))
	dec_cred := string(b_dec_cred)
	fmt.Printf(dec_cred)
	fmt.Printf("dbhost: " + string(myconf.DBHost))
	fmt.Printf("constring: " + string(myconf.DBHost) + string(myconf.DBPort) + "/" + string(myconf.Dbkicker))
	db, err := sql.Open("mysql", dec_cred+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/" + myconf.Dbkicker )
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := gin.Default()
	if err != nil {
		panic(err)
	}
	// router.LoadHTMLGlob("html/*")
	router.GET("/", func(c *gin.Context) {
		// c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	router.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	router.GET("/readiness", func(c *gin.Context) { c.JSON(http.StatusOK, "ok") })
	
	router.GET("/api/kicker", getKicker)
	router.POST("api/kicker/new", newKicker)
	router.GET("/api/kicker/:id", getKickerDetail)
	router.GET("/api/kicker/:id/latest", getKickerlatestGame)
	router.POST("/api/kicker/goal", KickerGoal)
	router.POST("/api/kicker/:id/:action", KickerPlayGame)
	router.GET("/api/games", getGames)
	router.GET("/api/games/:id/data", getGameData)
	// route for teams
	router.GET("/api/teams", ListTeams)
	router.GET("/api/teams/member", ListAllPlayer)
	router.POST("/api/teams/add", AddTeam)
	router.GET("/api/teams/:id/member", ListTeams)
	router.POST("/api/teams/:id/member/add", AddTeamMember)

	http.ListenAndServe(":8084", router)
	// router.Run(":8084")
}