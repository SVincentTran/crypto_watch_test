package apis

import (
	"cryptowatch/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEtherPriceHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	updates := helpers.GetUpdates()
	ctx.JSON(http.StatusOK, updates)
}
