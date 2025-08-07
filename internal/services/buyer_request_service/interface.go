package buyer_request_service

import "loopit/internal/models"

type BuyerRequestServiceInterface interface {
	CreateBuyerRequest(productID int, userCtx *models.UserContext) error
	UpdateBuyerRequestStatus(requestID int, updatedStatus string, userCtx *models.UserContext) error
	GetAllBuyerRequestsByStatus(productID int, status string) ([]models.BuyingRequest, error)
}
