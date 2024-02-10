package controllers

import (
	"goravel/app/http/requests"
	"goravel/app/models"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"message": "Goravel",
	})
}
func (r *UserController) Store(ctx http.Context) http.Response {
	var request requests.CreateUserRequest

	errors, err := ctx.Request().ValidateRequest(&request)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
	}
	if errors != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": errors.One(),
		})
	}

	password, err := facades.Hash().Make(request.Password)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
	}
	// if facades.Hash().Check(request.Password, password) {
	// 	// The passwords match...
	// 	fmt.Printf("password: %v\n", password)
	// }
	createUser := models.User{Username: request.Username, Password: password}
	if err := facades.Orm().Query().Create(&createUser); err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Json(http.StatusCreated, http.Json{
		"message": createUser,
	})
}
