package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Application Launched Success Fully")
}
