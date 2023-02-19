package delivery

import (
  "net/http"

  "github.com/gin-gonic/gin"
)


func Render(ctx *gin.Context) {
  ctx.HTML(http.StatusOK, "index.gohtml", gin.H{
    "Title": "Home",
  })
}
