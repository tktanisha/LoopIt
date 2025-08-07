package utils

import (
	"context"
	"fmt"
	"loopit/internal/config"
	"loopit/internal/constants"
	"loopit/internal/enums"
	"loopit/internal/models"
	"regexp"
	"strings"
)

func GetAuthenticatedUserFromContext(ctx context.Context) (*models.UserContext, bool) {
	userCtxRaw := ctx.Value(constants.UserCtxKey)
	userCtx, ok := userCtxRaw.(*models.UserContext)
	if !ok || userCtx == nil {
		fmt.Println(config.Red + "Unauthorized access. Please login first." + config.Reset)
		return nil, false
	}
	return userCtx, true
}

func ShowBanner() {
	fmt.Println()
	fmt.Println(config.Green + strings.Repeat("═", 80) + config.Reset)
	fmt.Println(config.Green + "║" + centerText("", 78) + "║" + config.Reset)
	fmt.Println(config.Green + "║" + centerText(config.Bold+"🚀 WELCOME TO Loop IT CLI PROJECT", 80) + "║" + config.Reset)
	fmt.Println(config.Green + "║" + centerText("", 78) + "║" + config.Reset)
	fmt.Println(config.Green + strings.Repeat("═", 80) + config.Reset)
	fmt.Println()
}

func PrintAllFeatures(role enums.Role) {
	fmt.Println(config.Cyan + "\nFEATURES" + config.Cyan)
	fmt.Println("1. Become a Lender")
	fmt.Println("2. Explore Products")
	fmt.Println("3. Get product by ID")
	fmt.Println("4. Create a new product")
	fmt.Println("5. Create a Buyer Request")
	fmt.Println("6. Get all Buyer Requests")
	fmt.Println("7. Update Buyer Request Status")
	fmt.Println("8. Get Order History")
	fmt.Println("9. Get Order History as Lender")
	fmt.Println("10. All orders that can be mark as returned")
	fmt.Println("11. Mark Order As Returned")
	fmt.Println("12. Create Return Request")
	fmt.Println("13. Update Return Request Status")
	fmt.Println("14. Logout")
	fmt.Println("15. Exit")
	fmt.Println()
}

// Helper to center text in given width
func centerText(text string, width int) string {
	padding := (width - len(stripAnsi(text))) / 2
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-padding-len(stripAnsi(text)))
}

// Helper to strip ANSI escape codes for correct padding
func stripAnsi(str string) string {
	ansiEscape := `\x1b\[[0-9;]*m`
	re := regexp.MustCompile(ansiEscape)
	return re.ReplaceAllString(str, "")
}
