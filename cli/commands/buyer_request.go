package commands

import (
	"fmt"
	"loopit/cli/utils"
	"loopit/internal/config"
	"loopit/internal/models"
	"os"

	"github.com/olekukonko/tablewriter"
)

// 1. Buy a Product (Create a Buyer Request)
func CreateBuyerRequest(userCtx *models.UserContext) {
	productID := utils.IntConversion(utils.Input("Enter Product ID to buy: "))

	err := BuyerRequestService.CreateBuyerRequest(productID, userCtx)
	if err != nil {
		fmt.Println(config.Red+"Error creating buyer request:"+config.Reset, err)
		return
	}

	fmt.Println(config.Green + "Buyer request created successfully!" + config.Reset)
}

// 2. Get all buyer requests (status = pending, approved, rejected)
func GetAllBuyerRequests() {
	productId := utils.IntConversion(utils.Input("Enter Product ID to fetch buyer requests: "))
	requests, err := BuyerRequestService.GetAllBuyerRequestsByStatus(productId, "pending")
	if err != nil {
		fmt.Println(config.Red+"Error fetching buyer requests:"+config.Reset, err)
		return
	}

	if len(requests) == 0 {
		fmt.Println(config.Yellow + "No buyer requests found." + config.Reset)
		return
	}

	fmt.Println("\nBuyer Requests:")

	table := tablewriter.NewWriter(os.Stdout)
	table.Header("ID", "Product ID", "Requested By", "Status", "Created At")

	for _, r := range requests {
		table.Append([]string{
			fmt.Sprintf("%d", r.ID),
			fmt.Sprintf("%d", r.ProductID),
			fmt.Sprintf("%d", r.RequestedBy),
			r.Status,
			r.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	table.Bulk(true)
	table.Render()
}

// 3. Update Buyer Request Status
func UpdateBuyerRequestStatus(userCtx *models.UserContext) {
	reqID := utils.IntConversion(utils.Input("Enter Buyer Request ID to update: "))

	statusOptions := []string{"approved", "rejected"}
	_, selectedStatus := utils.SelectFromList("Select new status", statusOptions)
	if selectedStatus == "" {
		fmt.Println(config.Red + "Status selection cancelled." + config.Reset)
		return
	}

	err := BuyerRequestService.UpdateBuyerRequestStatus(reqID, selectedStatus, userCtx)
	if err != nil {
		fmt.Println(config.Red+"Error updating status:"+config.Reset, err)
		return
	}

	fmt.Println(config.Green + "Buyer request status updated to '" + selectedStatus + "' successfully!" + config.Reset)
}
