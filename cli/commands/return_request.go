package commands

import (
	"fmt"
	"loopit/cli/utils"
	"loopit/internal/config"
	"loopit/internal/models"
)

// 1. Create Return Request (by lender)
func CreateReturnRequest(userCtx *models.UserContext) {
	orderID := utils.IntConversion(utils.Input("Enter Order ID to return: "))

	err := ReturnRequestService.CreateReturnRequest(userCtx.ID, orderID)
	if err != nil {
		fmt.Println(config.Red+"Error creating return request:"+config.Reset, err)
		return
	}

	fmt.Println(config.Green + "Return request created successfully!" + config.Reset)
}

// // 2. Get all pending Return Requests (for user)
// func GetAllPendingReturnRequests(userCtx *models.UserContext) {
// 	// You can update this to allow filtering by product/lender
// 	requests, err := ReturnRequestService.GetAllReturnRequestsByStatus("pending")
// 	if err != nil {
// 		fmt.Println(config.Red+"Error fetching return requests:"+config.Reset, err)
// 		return
// 	}

// 	userRequests := []models.ReturnRequest{}
// 	for _, req := range requests {
// 		order, _ := OrderService.GetOrderByID(req.OrderID)
// 		if order != nil && order.UserID == userCtx.ID {
// 			userRequests = append(userRequests, req)
// 		}
// 	}

// 	if len(userRequests) == 0 {
// 		fmt.Println(config.Yellow + "No pending return requests for your orders." + config.Reset)
// 		return
// 	}

// 	fmt.Println("\nPending Return Requests:")
// 	table := tablewriter.NewWriter(os.Stdout)
// 	table.Header("ID", "Order ID", "Requested By", "Status", "Created At")

// 	for _, r := range userRequests {
// 		table.Append([]string{
// 			fmt.Sprintf("%d", r.ID),
// 			fmt.Sprintf("%d", r.OrderID),
// 			fmt.Sprintf("%d", r.RequestedBy),
// 			r.Status,
// 			r.CreatedAt.Format(time.RFC822),
// 		})
// 	}
// 	table.Render()
// }

// 3. Update Return Request Status (accept/reject by user who placed the order)
func UpdateReturnRequestStatus(userCtx *models.UserContext) {
	reqID := utils.IntConversion(utils.Input("Enter Return Request ID to update: "))

	statusOptions := []string{"accepted", "rejected"}
	_, selectedStatus := utils.SelectFromList("Select new status", statusOptions)
	if selectedStatus == "" {
		fmt.Println(config.Red + "Status selection cancelled." + config.Reset)
		return
	}

	err := ReturnRequestService.UpdateReturnRequestStatus(userCtx.ID, reqID, selectedStatus)
	if err != nil {
		fmt.Println(config.Red+"Error updating return request status:"+config.Reset, err)
		return
	}

	fmt.Println(config.Green + "Return request status updated to '" + selectedStatus + "' successfully!" + config.Reset)
}
