package cli

import (
	"context"
	"fmt"
	"loopit/cli/commands"
	"loopit/cli/utils"
	"loopit/internal/config"
	"strings"
)

func FeatureMenu(ctx context.Context) {

	userCtx, ok := utils.GetAuthenticatedUserFromContext(ctx)
	if !ok || userCtx == nil {
		fmt.Println(config.Red + "Unauthorized access. Please login first." + config.Reset)
		return
	}

	for {
		utils.PrintAllFeatures(userCtx.Role)
		fmt.Print(config.Yellow + "Choose an option: " + config.Reset)
		var choice string
		fmt.Scanln(&choice)

		switch strings.TrimSpace(choice) {
		case "1":
			commands.BecomeLender(userCtx)
		case "2":
			commands.GetAllProducts()
		case "3":
			commands.GetProductByID()
		case "4":
			commands.CreateProduct(userCtx)
		case "5":
			commands.CreateBuyerRequest(userCtx)
		case "6":
			commands.GetAllBuyerRequests()
		case "7":
			commands.UpdateBuyerRequestStatus(userCtx)
		case "8":
			commands.GetOrderHistory(userCtx)
		case "9":
			commands.GetLenderOrders(userCtx)
		case "10":
			commands.GetAllApprovedAwaitingOrders(userCtx)
		case "11":
			commands.MarkOrderAsReturned(userCtx)
		case "12":
			commands.CreateReturnRequest(userCtx)
		case "13":
			commands.UpdateReturnRequestStatus(userCtx)
		case "14":
			commands.AuthLogout(&ctx)
			return
		case "15":
			fmt.Println("Exiting. Goodbye!")
			return
		default:
			fmt.Println(config.Red + "Invalid option. Try again." + config.Reset)
		}
	}
}
