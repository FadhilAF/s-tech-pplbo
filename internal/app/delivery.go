package app

import (
  user_delivery "github.com/FadhilAF/s-tech-pplbo/internal/delivery/user"
)

type deliveries struct {
  user user_delivery.UserDelivery
}

func (app *App) initDelivery() {
	var deliveries deliveries
  deliveries.user = user_delivery.NewUserDelivery(app.usecase.user)

	app.delivery = deliveries
}
