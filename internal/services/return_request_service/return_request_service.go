package return_request_service

import (
	"errors"
	"loopit/internal/models"
	"loopit/internal/repository/order_repo"
	"loopit/internal/repository/product_repo"
	"loopit/internal/repository/return_request_repo"
	"time"
)

type returnRequestService struct {
	orderRepo         order_repo.OrderRepo
	productRepo       product_repo.ProductRepo
	returnRequestRepo return_request_repo.ReturnRequestRepo
}

func NewReturnRequestService(orderRepo order_repo.OrderRepo, productRepo product_repo.ProductRepo, rrRepo return_request_repo.ReturnRequestRepo) ReturnRequestServiceInterface {
	return &returnRequestService{
		orderRepo:         orderRepo,
		returnRequestRepo: rrRepo,
		productRepo:       productRepo,
	}
}

func (s *returnRequestService) CreateReturnRequest(userID int, orderID int) error {
	order, err := s.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return err
	}

	// Must be "in_use"
	if order.Status != "in_use" {
		return errors.New("order is not in 'in_use' status")
	}

	productID := order.ProductID
	product, err := s.productRepo.FindByID(productID)

	if err != nil {
		return err
	}

	if product.Product.LenderID != userID {
		return errors.New("user is not lender of the order's product")
	}

	if order.Status != "in_use" {
		return errors.New("order is not in 'in_use' status and cannot be returned")
	}

	returnRequest := models.ReturnRequest{
		OrderID:     orderID,
		RequestedBy: userID,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	if err := s.orderRepo.UpdateOrderStatus(orderID, "returned_requested"); err != nil {
		return err
	}

	return s.returnRequestRepo.CreateReturnRequest(returnRequest)
}

func (s *returnRequestService) UpdateReturnRequestStatus(userID int, reqID int, newStatus string) error {
	// Only accept or reject allowed
	if newStatus != "accepted" && newStatus != "rejected" {
		return errors.New("invalid status update")
	}

	req, err := s.returnRequestRepo.GetReturnRequestByID(reqID)
	if err != nil {
		return err
	}

	if req.Status != "pending" {
		return errors.New("return request is not in pending status")
	}

	order, err := s.orderRepo.GetOrderByID(req.OrderID)
	if err != nil {
		return err
	}

	if order.UserID != userID {
		return errors.New("user does not own this order")
	}
	return s.returnRequestRepo.UpdateReturnRequestStatus(req.ID, newStatus)
}

// func (s *returnRequestService) GetPendingReturnRequests(userID int) ([]models.ReturnRequest, error) {
// 	allRequests, _ := s.returnRequestRepo.GetAllReturnRequests([]string{"pending"})
// 	var userRequests []models.ReturnRequest
// 	for _, req := range allRequests {
// 		if req.LenderID == userID {
// 			userRequests = append(userRequests, req)
// 		}
// 	}
// 	return userRequests, nil
// }
