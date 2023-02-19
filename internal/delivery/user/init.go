package delivery

import (
	usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserDelivery interface {
	CreateUser(ctx *gin.Context)
	// DeleteUser(ctx *gin.Context)
	// GetUser(ctx *gin.Context)
	// ListUser(ctx *gin.Context)
	// UpdateUser(ctx *gin.Context)
}

var _ UserDelivery = &userHandler{}

func NewUserDelivery(usecase usecase.UserUsecase) UserDelivery {
	return &userHandler{
		usecase: usecase,
	}
}

type userHandler struct {
	usecase usecase.UserUsecase
}
