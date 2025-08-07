package controllers

// import (
// 	"loopit/internal/models"
// 	service "loopit/internal/services"
// )

// type AuthController struct {
// 	authService service.AuthService
// }

// // func NewAuthController() *AuthController {
// // 	return &AuthController{
// // 		authService: service.NewAuthService(),
// // 	}
// // }

// func (a *AuthController) Register(name, email, password string) (*models.User, error) {
// 	err := a.authService.Register(name, email, password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &models.User{FullName: name, Email: email}, nil
// }

// func (a *AuthController) Login(email, password string) (*models.User, error) {
// 	user, err := a.authService.Login(email, password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }
