package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationStatus struct {
	Status bool `json:"status"`
}

func GetLoginStatus(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	authRes := AuthenticationStatus{
		Status: true,
	}

	ctx.JSON(http.StatusOK, authRes)
}
