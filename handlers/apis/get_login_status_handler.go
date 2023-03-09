package apis

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthenticationStatus struct {
	Status bool   `json:"status"`
	UserId string `json:"user_id"`
}

func GetLoginStatus(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	session := sessions.Default(ctx)
	profile := session.Get("profile")

	userId := ""
	if profile != nil {
		convertedProfile := profile.(map[string]interface{})
		userId = convertedProfile["nickname"].(string)
	}

	authRes := AuthenticationStatus{
		Status: true,
		UserId: userId,
	}

	ctx.JSON(http.StatusOK, authRes)
}
