package app

import (
	"database/sql"
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/FadhilAF/s-tech-pplbo/config"

	"github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type App struct {
	Config   config.Config
	delivery deliveries
	usecase  usecases
	store    repository.Store
}

func (app *App) createHandlers() *gin.Engine {
	router := gin.Default()

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowHeaders = append(corsCfg.AllowHeaders, "Accept")

	corsCfg.AllowAllOrigins = true

	router.Use(cors.New(corsCfg))

	viewRouterGroup := router.Group("/")
	apiRouterGroup := router.Group("/api")

	app.viewHandler(viewRouterGroup)
	app.apiHandler(apiRouterGroup)

	// mesin untuk menampilkan semua routes yang ada?
	routes := router.Routes()
	if gin.Mode() == gin.DebugMode {
		fmt.Println()
		for _, v := range routes {
			path := color.InBold(v.Path)
			method := color.InYellow(fmt.Sprintf("%-6s", v.Method))
			fmt.Println(method, path)
		}
		fmt.Println()
	}
	// mesin untuk menampilkan semua routes yang ada? selesai

	return router
}

func (app *App) StartServer() {

}

func New(config config.Config, db *sql.DB) App {
	app := App{}

	app.Config = config

	app.store = repository.NewPostgresStore(db)
	app.initUsecase()
	app.initDelivery()

	return app
}
