package customer

import (
	. "neema.co.za/rest/modules/customer/internal/api"
)

func GetModule() *Module {
	api := BuildApi()
	handleRoutes(api)
	module := Module(*api) //Module is an alias of Api
	return &module
}

func handleRoutes(api *Api) {
	api.Get("/:id", api.GetCustomerByID)
	api.Post("", api.GetCustomerByID)
}
