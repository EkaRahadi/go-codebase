package utils

import (
	"reflect"
	"strings"
	"unicode"

	apperror "github.com/EkaRahadi/go-codebase/internal/error"

	"github.com/EkaRahadi/go-codebase/internal/dto"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type (
	Validator interface {
		Validate(obj interface{}) error
	}

	customValidator struct {
		validate   *validator.Validate
		translator ut.Translator
	}
)

func NewCustomValidator() *customValidator {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	//nolint:errcheck // no need to check
	validate.RegisterValidation("username", validateUsername)
	//nolint:errcheck // no need to check
	validate.RegisterValidation("password", validatePassword)
	//nolint:errcheck // no need to check
	en_translations.RegisterDefaultTranslations(validate, trans)
	errorStructTranslator := &customValidator{
		validate:   validate,
		translator: trans,
	}

	return errorStructTranslator
}

func (vldtr customValidator) Validate(obj interface{}) error {
	err := vldtr.validate.Struct(obj)
	errsData := make([]dto.ValidationErrorResponse, 0)
	if err != nil {
		//nolint:errcheck // no need to check
		validatorErrs := err.(validator.ValidationErrors)
		for _, e := range validatorErrs {
			translatedErr := e.Translate(vldtr.translator)
			if e.ActualTag() == "username" {
				translatedErr = "Username must contain only string and digit"
			} else if e.ActualTag() == "password" {
				translatedErr = "Password must have at least 8 characters, 1 symbol, 1 capital letter, and 1 number"
			}

			errData := dto.ValidationErrorResponse{
				Field:   e.Field(),
				Message: translatedErr,
			}
			errsData = append(errsData, errData)
		}
	}

	if len(errsData) > 0 {
		validationErr := apperror.NewValidationError("input validation error", errsData)
		return validationErr
	}

	return nil
}

func validatePassword(fl validator.FieldLevel) bool {
	// Requires at least 8 characters, 1 symbol, 1 capital letter, and 1 number
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	var hasSymbol, hasUpper, hasLower, hasDigit bool

	for _, char := range password {
		switch {
		case unicode.IsSymbol(char), unicode.IsPunct(char):
			hasSymbol = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	return hasSymbol && hasUpper && hasLower && hasDigit
}

func validateUsername(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var validUsername bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			validUsername = true
		case unicode.IsLower(char):
			validUsername = true
		case unicode.IsDigit(char):
			validUsername = true
		default:
			return false
		}
	}

	return validUsername
}
