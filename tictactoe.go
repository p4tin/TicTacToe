package main

import (
	"net/http"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"github.com/p4tin/tictactoe/game"
)

func main() {
	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("tictactoe", store))
	//options := sessions.Options{MaxAge: 20}
	//store.Options(options)

	//router.Use(func(c *gin.Context) {
	//	switch(c.Request.URL.Path) {
	//	case "/start":
	//		fallthrough
	//	case "/move":
	//		fallthrough
	//	case "/board":
	//		session := sessions.Default(c)
	//		username := session.Get("user")
	//		if username == nil {
	//			session.Clear()
	//			session.Save()
	//			c.Redirect(http.StatusTemporaryRedirect, "/login")
	//			return
	//		} else {
	//			//t := time.Now()
	//			//expires := session.Get("expire")
	//			//if t >= expires {
	//			//	session.Clear()
	//			//	session.Save()
	//			//	c.Redirect(http.StatusTemporaryRedirect, "/login")
	//			//	return
	//			//}
	//		}
	//		co, err := c.Cookie("tictactoe")
	//		if err != nil {
	//			log.Println("||| Error - ", err, "|||")
	//		} else {
	//			log.Println("||| Cookie - ", co, "|||")
	//		}
	//	}
	//
	//})

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", home)

	router.GET("/login", loginForm)
	router.POST("/login", login)
	router.GET("/logout", logout)

	router.GET("/register", registerForm)
	router.POST("/register", register)

	router.GET("/board", displayBoard)
	router.POST("/move", playerMove)
	router.POST("/start", startGame)

	router.Run()
}



func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home",
	})
}

func displayBoard(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if username == nil {
		session.Clear()
		session.Save()
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	user := game.Players[username.(string)]
	c.HTML(http.StatusOK, "board.html", gin.H{
		"title": "Board", "user": user, "ingame": game.Players[username.(string)].Ingame,
	})
}

func loginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if pl, ok := game.Players[username]; ok {
		if pl.Password == password {
			fmt.Printf("Login - Save Sessions Info: %+V\n%+v\n", game.Players, pl)
			session := sessions.Default(c)
			session.Set("user", pl.Name)
//			session.Set("expires", time.Now())
			session.Save()
			c.HTML(http.StatusOK, "board.html", gin.H{
				"title": "Login", "user": pl,
			})
		} else {
			c.HTML(http.StatusNotFound, "login.html", gin.H{
				"title": "Login",
			})
		}
	} else {
		c.HTML(http.StatusNotFound, "login.html", gin.H{
			"title": "Login",
		})
	}
}


func registerForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println("Trying to register:", username, "with password:", password)
	pl, err := game.NewPlayer(username, password)
	if err != "" {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register", "error": err,
		})
	}
	session := sessions.Default(c)
	session.Set("user", pl.Name)
	session.Save()
	c.HTML(http.StatusOK, "board.html", gin.H{
		"title": "Login", "user": pl,
	})
}

func playerMove(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user").(string)
	if username == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	square := c.PostForm("moveButton")
	game.PlayerMove(square, username)

	c.Redirect(http.StatusMovedPermanently, "/board")
}

func startGame(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user").(string)
	if username == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		return
	}
	game.StartGame(username)
	c.Redirect(http.StatusMovedPermanently, "/board")
}

