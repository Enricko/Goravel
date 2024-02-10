package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateUserRequest struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (r *CreateUserRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"username": "required|max_len:255",
		"password": "required|max_len:255",
	}
}

func (r *CreateUserRequest) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateUserRequest) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
