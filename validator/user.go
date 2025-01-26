package validator

import (
	"next-learn-go-sqlc/infrastructure/database/sqlc"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserValidator interface {
	UserValidate(user sqlc.User) error
}

type userValidator struct{}

func NewUserValidator() UserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user sqlc.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}
