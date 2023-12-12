package paymentreceive

import (
	. "neema.co.za/rest/modules/paymentreceive/internal/api"
)

func GetModule() *Module {
	api := BuildApi()
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {	
	api.Get("/", api.GetAllPaymentReceived)
	api.Get("/:id", api.GetPaymentReceivedByID)
	api.Delete("/:id", api.DeletePaymentReceived)
	api.Post("/", api.CreatePaymentReceived)
}
