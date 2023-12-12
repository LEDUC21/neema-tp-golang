package service

import (
	"neema.co.za/rest/utils/models"
	"neema.co.za/rest/utils/resources"
)

func (s *Service) GetPaymentReceivedByID(id int) (*resources.PaymentReceivedResource, error) {
	// Business logic (if any)
	return s.Repository.GetPaymentReceivedByID(id)
}

func (s *Service) GetAllPaymentReceived() ([]resources.PaymentReceivedResource, error) {
	// Business logic (if any)
	return s.Repository.GetAllPaymentReceived()
}

func (s *Service) GetByIdDeletePaymentReceived(id int) (models.PaymentReceived, error) {
	// Business logic (if any)
	return s.Repository.GetByIdDeletePaymentReceived(id)
}

func (s *Service) DeletePaymentReceived(id int) (models.PaymentReceived, error) {
	// Business logic (if any)
	return s.Repository.DeletePaymentReceived(id)
}

func (s *Service) CreatePaymentReceived(payment_received models.PaymentReceived) (models.PaymentReceived, error) {
	// Business logic (if any)
	return s.Repository.CreatePaymentReceived(payment_received)
}
