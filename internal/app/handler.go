package app

import (
	"github.com/FadhilAF/s-tech-pplbo/internal/route"
	"github.com/gin-gonic/gin"
)

func (app *App) viewHandler(router *gin.RouterGroup) {

}

func (app *App) apiHandler(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	route.UserRoutes(userGroup, app.delivery.user)
}
