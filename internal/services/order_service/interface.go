package order_service

import "loopit/internal/models"

type OrderServiceInterface interface {
	// CreateOrder(order *models.Order) error-//TODO -need nhi h kyunki buy request se order create ho rha
	UpdateOrderStatus(orderID int, newStatus string) error
	GetOrderHistory(userCtx *models.UserContext, filterStatus []string) ([]*models.Order, error)
	GetAllApprovedAwaitingOrders(userCtx *models.UserContext) ([]*models.Order, error) // for Lender to get all orders that are returned and awaiting status
	MarkOrderAsReturned(orderID int, userCtx *models.UserContext) error                // Lender marks the product as returned after receiving it
	GetLenderOrders(userCtx *models.UserContext) ([]*models.Order, error)              // Lender gets all orders where they are the lender

	// // Return-related
	// SendReturnRequest(orderID int, reason string) error        // Lender requests return
	// ApproveReturnRequest(orderID int) error                    // User approves return

}
