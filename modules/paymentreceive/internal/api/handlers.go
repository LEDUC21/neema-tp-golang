package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	. "neema.co.za/rest/utils/resources"
)

func (this *Api) GetPaymentReceivedByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	payment_received, err := this.Service.GetPaymentReceivedByID(id)
	if err != nil {
		SendErrors("PaymentReceived don't exist", c)
		return nil
	}

	if payment_received.ID == 0 {
		SendErrors("PaymentReceived list is empty", c)
		return nil
	}

	SendResponses(payment_received, "Payment_received successfuly recovered", c, "ok")
	return nil
}

func (this *Api) GetAllPaymentReceived(c *fiber.Ctx) error {
	payment_received, err := this.Service.GetAllPaymentReceived()
	if err != nil {
		SendErrors("PaymentReceived don't exist", c)
		return nil
	}

	SendResponse(payment_received, "Get All Successfuly", c, "ok")
	return nil
}

func (this *Api) DeletePaymentReceived(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	payment_received, err := this.Service.GetByIdDeletePaymentReceived(id)
	if err != nil {
		SendErrors("Payment_received don't exist for deletion", c)
		return nil
	}

	if payment_received.Used_amount == "$0.00" {

		payment_received, err := this.Service.DeletePaymentReceived(id)
		if err != nil {
			SendError(err, c)
			return nil
		}
		SendResponses(payment_received, "Payment_received deleted successfuly", c, "ok")

	} else {
		SendErrors("Deletion impossible, payment used", c)
		return nil
	}

	return nil
}

func (this *Api) CreatePaymentReceived(c *fiber.Ctx) error {
	payment_received, err := this.Service.ParsePaymentReceivedVariable(c)
	if err != nil {
		SendError(err, c)
		return nil
	}

	payment_receiveds, errs := this.Service.CreatePaymentReceived(payment_received)
	if errs != nil {
		SendError(err, c)
		return nil
	}

	SendResponses(payment_receiveds, "Payment_received create Successfuly", c, "create")
	
	return nil
}
