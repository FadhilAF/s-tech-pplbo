package route

import (
  delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/user"

  "github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, handler delivery.UserDelivery) {
  router.POST("/", handler.CreateUser)
  // router.GET("/", handler.GetAllUser)
  // router.GET("/:id", handler.GetUserById)
  // router.PUT("/:id", handler.UpdateUser)
  // router.DELETE("/:id", handler.DeleteUser)
}
