package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func KickerStartgame(c *gin.Context) {
	name := c.Param("name")
	message := "comming soon! " + name + " action: StartGame"
	c.String(http.StatusOK, message)

}

func KickerGoal(c *gin.Context) {
	name := c.Param("name")
	message := "comming soon! " + name + " action: Goal"
	c.String(http.StatusOK, message)

}

func KickerEndgame(c *gin.Context) {
	name := c.Param("name")
	message := "comming soon! " + name + " action: StartGame"
	c.String(http.StatusOK, message)

}
