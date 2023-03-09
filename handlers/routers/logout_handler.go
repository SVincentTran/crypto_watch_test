package routers

import (
	"cryptowatch/config"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our logout.
func LogoutHandler(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logoutUrl, err := url.Parse("https://" + cfg.Auth0.Domain + "/v2/logout")
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		scheme := "http"
		if ctx.Request.TLS != nil {
			scheme = "https"
		}

		returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		parameters := url.Values{}
		parameters.Add("returnTo", returnTo.String())
		parameters.Add("client_id", cfg.Auth0.ClientId)
		logoutUrl.RawQuery = parameters.Encode()

		session := sessions.Default(ctx)
		session.Clear()
		session.Save()

		log.Printf("Logout Detail %v", logoutUrl.String())

		ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
	}
}
