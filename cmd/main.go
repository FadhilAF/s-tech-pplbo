package main

import (
	"log"

	httpCommon "github.com/FadhilAF/s-tech-pplbo/common/http"

	// "github.com/FadhilAF/s-tech-pplbo/config"

	landingPageDelivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/landing"
)

func main() {
	httpServer := httpCommon.NewHTTPServer()

	httpServer.Router.LoadHTMLGlob("internal/template/*.gohtml")

	httpServer.Router.GET("/", landingPageDelivery.Render)

	// appConfig := config.LoadConfig(".env")

	log.Fatalln(httpServer.Router.Run()) //kalo print, keluar di terminal, kalo Sprint return string
}
