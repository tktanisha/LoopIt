package buyer_request_service

import (
	"errors"
	"loopit/internal/enums"
	"loopit/internal/models"
	"loopit/internal/repository/buyer_request_repo"
	"loopit/internal/repository/category_repo"
	"loopit/internal/repository/order_repo"
	"loopit/internal/repository/product_repo"
	"time"
)

type BuyerRequestService struct {
	buyerRequestRepo buyer_request_repo.BuyerRequestRepo
	productRepo      product_repo.ProductRepo
	orderRepo        order_repo.OrderRepo
	categoryRepo     category_repo.CategoryRepo
}

func NewBuyerRequestService(
	buyerReqRepo buyer_request_repo.BuyerRequestRepo,
	productRepo product_repo.ProductRepo,
	orderRepo order_repo.OrderRepo,
	categoryRepo category_repo.CategoryRepo,
) BuyerRequestServiceInterface {
	return &BuyerRequestService{
		buyerRequestRepo: buyerReqRepo,
		productRepo:      productRepo,
		orderRepo:        orderRepo,
		categoryRepo:     categoryRepo,
	}
}

func (s *BuyerRequestService) CreateBuyerRequest(productID int, userCtx *models.UserContext) error {
	product, err := s.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("product not found")
	}
	if !product.Product.IsAvailable {
		return errors.New("product not available")
	}

	// Check if a pending or approved request already exists
	allRequests, err := s.buyerRequestRepo.GetAllBuyerRequests([]string{"pending", "approved"})
	if err != nil {
		return err
	}
	for _, req := range allRequests {
		if req.ProductID == productID && req.RequestedBy == userCtx.ID {
			return errors.New("a pending or approved request already exists")
		}
	}

	newRequest := models.BuyingRequest{
		ProductID:   productID,
		RequestedBy: userCtx.ID,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	if err := s.buyerRequestRepo.CreateBuyerRequest(newRequest); err != nil {
		return err
	}

	return nil
}

func (s *BuyerRequestService) UpdateBuyerRequestStatus(requestID int, updatedStatus string, userCtx *models.UserContext) error {
	if userCtx.Role != enums.RoleLender {
		return errors.New("unauthorized: only lenders can update request status")
	}

	if updatedStatus != "approved" && updatedStatus != "rejected" {
		return errors.New("invalid status: only 'approved' or 'rejected' allowed")
	}

	allRequests, err := s.buyerRequestRepo.GetAllBuyerRequests(nil)
	if err != nil {
		return err
	}

	var req *models.BuyingRequest
	for i := range allRequests {
		if allRequests[i].ID == requestID {
			req = &allRequests[i]
			break
		}
	}
	if req == nil {
		return errors.New("buyer request not found")
	}

	// If rejected: just update status
	if updatedStatus == "rejected" {
		if err := s.buyerRequestRepo.UpdateStatusBuyerRequest(requestID, "rejected"); err != nil {
			return err
		}
		return nil
	}

	// If approved: update status and create new order
	product, err := s.productRepo.FindByID(req.ProductID)
	if err != nil {
		return errors.New("product not found")
	}

	category, err := s.categoryRepo.FindByID(product.Category.ID)
	if err != nil {
		return errors.New("category not found")
	}

	newOrder := models.Order{
		ProductID:      req.ProductID,
		UserID:         req.RequestedBy,
		StartDate:      time.Now(),
		EndDate:        time.Time{},
		TotalAmount:    category.Price,
		SecurityAmount: category.Security,
		Status:         "in_use",
		CreatedAt:      time.Now(),
	}

	if err := s.orderRepo.CreateOrder(newOrder); err != nil {
		return err
	}

	if err := s.buyerRequestRepo.UpdateStatusBuyerRequest(requestID, "approved"); err != nil {
		return err
	}
	return nil
}

func (s *BuyerRequestService) GetAllBuyerRequestsByStatus(productID int, status string) ([]models.BuyingRequest, error) {
	filtered, err := s.buyerRequestRepo.GetAllBuyerRequests([]string{status})
	if err != nil {
		return nil, err
	}

	result := []models.BuyingRequest{}
	for _, req := range filtered {
		if req.ProductID == productID {
			result = append(result, req)
		}
	}

	return result, nil
}
