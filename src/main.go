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

type games struct {
	GAMEID    int    `json:"gameid"`
	KICKERID  int    `json:"kickerid"`
	STARTDATE string `json:"startdate"`
	ENDDATE   string `json:"enddate"`
}

type gamedata struct {
	GDID        int     `json:"gdid"`
	GAMEID      int     `json:"gameid"`
	KICEKRID    int     `json:"kickerid"`
	GOALDATE    string  `json:"goaldate"`
	TEAM_A_GOAL int     `json:"team_a_goal"`
	TEAM_B_GOAL int     `json:"team_b_goal"`
	HUMIDITY    float32 `json:"humidity"`
	TEMPERATURE float32 `json:"temperature"`
	TEAM_A_ID   int     `json:"team_a_id"`
	TEAM_B_ID   int     `json:"team_b_id"`
}
type conf struct {
	DBHost     string `yaml:"dbhost"`
	DBPort     string `yaml:"dbport"`
	Credential string `yaml:"dbcredential"`
}

func (c *conf) GetConfig() *conf {
	confContent, err := ioutil.ReadFile("config.yml")
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

	db, err := sql.Open("mysql", dec_cred+"@tcp("+myconf.DBHost+":"+myconf.DBPort+")/smartkicker")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := gin.Default()
	if err != nil {
		panic(err)
	}
	router.LoadHTMLGlob("html/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	router.GET("/api/kicker", getKicker)
	router.POST("api/kicker/new", newKicker)
	router.GET("/api/kicker/:name", getKickerDetail)
	router.GET("/api/kicker/:name/goal", KickerGoal)
	router.GET("/api/kicker/:name/startgame", KickerStartgame)
	router.GET("/api/kicker/:name/endgame", KickerEndgame)
	router.GET("/api/games", getGames)
	router.Run(":8083")
}
