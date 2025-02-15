package app

import (
	"banking/client"
	"banking/controller"
)

type Presenter struct {
	UserPresenter *controller.Controller
}

func InitFactory() Presenter {
	userClient := client.NewBanking()
	userController := controller.NewController(userClient)

	return Presenter{
		UserPresenter: userController,
	}
}
