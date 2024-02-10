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
	header := ctx.Request().Header("Authorization","")
	payload, err := facades.Auth().Parse(ctx, header)
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": payload,
	})
}

func (r *UserController) Logout(ctx http.Context) http.Response {
	header := ctx.Request().Header("Authorization","")
	payload, err := facades.Auth().Parse(ctx, header)
	err = facades.Auth().Logout(ctx)
	
	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"message": err.Error(),
		})
	}

	return ctx.Response().Success().Json(http.Json{
		"message": payload,
	})
}
func (r *UserController) Login(ctx http.Context) http.Response {
	var request requests.LoginUserRequest

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

	var existingUser models.User
	if err := facades.Orm().Query().Model(&existingUser).Where("username = ?", request.Username).First(&existingUser); err == nil {
		if err != nil {
			return ctx.Response().Json(http.StatusBadRequest,http.Json{
				"message": err.Error(),
			})
		}
		if facades.Hash().Check(request.Password, existingUser.Password) {
			token, err := facades.Auth().Login(ctx, &existingUser)
			if err != nil {
				return ctx.Response().Json(http.StatusBadRequest,http.Json{
					"message": err.Error(),
				})
			}
			return ctx.Response().Json(http.StatusAccepted,http.Json{
				"message": "Username already exists",
				"token":   token,
				"user":    existingUser,
			})
		}
	}
	return ctx.Response().Json(http.StatusAccepted,http.Json{
		"message": "asdasd",
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
