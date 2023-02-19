package http

import ( 
  "github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Router *gin.Engine
}

func NewHTTPServer() HTTPServer {
	router := gin.Default()

	return HTTPServer{
		Router: router,
	}
}
