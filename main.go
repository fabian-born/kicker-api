package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}



func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
    db, err := sql.Open("mysql", "root:root@tcp(192.168.69.22:8989)/smartkicker")
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
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbums)
    router.GET("/api/kicker", getKicker)
    router.GET("/api/kicker/:id", getKickerDetail)
    router.GET("/api/kicker/:id/*action", actionKicker)
    router.GET("/api/games", getGames)
    router.Run(":8083")
}
