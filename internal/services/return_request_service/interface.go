package return_request_service

type ReturnRequestServiceInterface interface {
	CreateReturnRequest(userID int, orderID int) error
	UpdateReturnRequestStatus(userID int, reqID int, newStatus string) error
	// GetPendingReturnRequests(userID int) ([]models.ReturnRequest, error)
}
