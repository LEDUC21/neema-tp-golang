//go:build wireinject

package paymentreceive

import (
	"github.com/google/wire"

	"neema.co.za/rest/modules/paymentreceive/internal/api"
	"neema.co.za/rest/modules/paymentreceive/internal/repository"
	"neema.co.za/rest/modules/paymentreceive/internal/service"
	"neema.co.za/rest/utils/app"
	"neema.co.za/rest/utils/database"
)

// New api handler
func BuildApi() *api.Api {
	panic(wire.Build(
		database.GetDatabase,
		app.NewFiberApp,
		repository.NewRepository,
		service.NewService,
		api.NewApi,
	))
}
