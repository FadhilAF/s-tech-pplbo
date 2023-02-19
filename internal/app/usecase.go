package app

import (
  user_usecase "github.com/FadhilAF/s-tech-pplbo/internal/usecase/user"
)

type usecases struct {
  user user_usecase.UserUsecase
}

func (app *App) initUsecase() {
	var usecases usecases
  usecases.user = user_usecase.NewUserUsecase(app.store)

	app.usecase = usecases
}
