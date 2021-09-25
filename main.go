package main

import (
	"chatroom/auth"
	"chatroom/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()
	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", auth.AuthenticateHeader, func(c *gin.Context) {
		// uncomment if testing tools support websocket upgrader
		//roomId := c.Param("roomId")
		//clientid := c.MustGet("clientid").(string)
		//serveWs(c.Writer, c.Request, roomId, clientid)
		c.HTML(200, "index.html", nil)
	})

	router.DELETE("/room/:roomId", auth.AuthenticateHeader, auth.CheckAdmin, func(c *gin.Context) {
		roomId := c.Param("roomId")
		status := deleteRooms(roomId)
		if status {
			c.JSON(http.StatusOK, "{\"message\":\"Deleted\"}")
		} else {
			c.JSON(http.StatusGone, "{\"message\":\"Deletion Failed\"}")
		}
	})

	router.GET("/user/room/:roomId", auth.AuthenticateHeader, auth.CheckAdmin, func(c *gin.Context) {
		roomId := c.Param("roomId")
		c.JSON(http.StatusOK, listUser(roomId))
	})

	router.DELETE("/user/room/:roomId/:username", auth.AuthenticateHeader, auth.CheckAdmin, func(c *gin.Context) {
		roomId := c.Param("roomId")
		username := c.Param("username")
		status := removeUser(roomId, username)
		if status {
			c.JSON(http.StatusOK, "{\"message\":\"Deleted\"}")
		} else {
			c.JSON(http.StatusGone, "{\"message\":\"Deletion Failed\"}")
		}
	})

	router.GET("/token", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		req := database.LoginRequest{
			Username: username,
			Password: password,
		}
		logistatus, adminstatus := database.Login(&req)
		if (username == "admin" && adminstatus == false) || logistatus == false {
			c.JSON(http.StatusUnauthorized, "{\"message\":\"Incorrect Username and Password\"}")
			c.Abort()
		} else {
			c.Next()
		}

	}, func(c *gin.Context) {
		username := c.Param("username")
		token, err := auth.CreateToken(username)
		if err == nil {
			c.JSON(http.StatusAccepted, token)
		} else {
			c.JSON(http.StatusInternalServerError, err)
		}
	})

	router.GET("/rooms", auth.AuthenticateHeader, func(c *gin.Context) {
		rooms := listRooms()
		c.JSON(http.StatusOK, rooms)
	})

	router.GET("/ws/:roomId", auth.AuthenticateHeader, func(c *gin.Context) {
		roomId := c.Param("roomId")
		clientid := c.MustGet("clientid").(string)
		serveWs(c.Writer, c.Request, roomId, clientid)
	})

	router.Run("0.0.0.0:8080")
}
