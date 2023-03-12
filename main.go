package main

import (
	"cryptowatch/authenticator"
	"cryptowatch/config"
	apiHandlers "cryptowatch/handlers/apis"
	"cryptowatch/handlers/middleware"
	routerHandlers "cryptowatch/handlers/routers"
	"cryptowatch/helpers"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()

	auth, err := authenticator.New(config)
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	db, err := helpers.InitDBConnection(&config.Postgres)
	if err != nil {
		log.Printf("Error while initializing DB: %v", err)
	}

	gob.Register(map[string]interface{}{})

	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	router.GET("/login", routerHandlers.LoginHandler(auth))
	router.GET("/auth-callback", routerHandlers.AuthenticationHandler(auth, db))
	router.GET("/logout", middleware.IsAuthenticated, routerHandlers.LogoutHandler(config))

	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	api.GET("/watch", middleware.IsAuthenticated, apiHandlers.GetEtherPriceHandler)
	api.GET("/login-status", middleware.IsAuthenticated, apiHandlers.GetLoginStatus)

	// Running a goroutine to handle the web socket subscribtion
	go helpers.InitWebSocketClient()
	router.Run(":3000")
}
