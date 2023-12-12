package repository

import (
	"github.com/gofiber/fiber/v2"
	. "neema.co.za/rest/utils/models"
	. "neema.co.za/rest/utils/resources"
)

func (r *Repository) GetPaymentReceivedByID(id int) (*PaymentReceivedResource, error) {
	var payment_received PaymentReceivedResource
	r.Table("payment_received").
			 Join("INNER", "customer", "payment_received.id_customer = customer.id").
    		 Where("payment_received.id = ?", id).
    		 Find(&payment_received)
		
	return &payment_received, nil
}

func (r *Repository) GetAllPaymentReceived() ([]PaymentReceivedResource, error) {
	var payment_received []PaymentReceivedResource
	r.Table("payment_received").
    		 Join("INNER", "customer", "payment_received.id_customer = customer.id").
    		 Where("payment_received.tag= ?", "1").
    		 Find(&payment_received)
		
	return payment_received, nil
}

func (r *Repository) GetByIdDeletePaymentReceived(id int) (PaymentReceived, error) {
	var payment_received PaymentReceived
	r.ID(id).Get(&payment_received)
		
	return payment_received, nil
}

func (r *Repository) DeletePaymentReceived(id int) (PaymentReceived, error) {
	var payment_received PaymentReceived
	r.ID(id).Delete(&payment_received)
		
	return payment_received, nil
}

func (r *Repository) ParsePaymentReceivedVariable(c *fiber.Ctx) (PaymentReceived, error) {
	var payment_received PaymentReceived
	c.BodyParser(&payment_received)
		
	return payment_received, nil
}

func (r *Repository) CreatePaymentReceived(payment_received PaymentReceived) (PaymentReceived, error) {
	r.Insert(&payment_received)
		
	return payment_received, nil
}
