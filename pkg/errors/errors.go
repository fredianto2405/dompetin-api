package errors

import (
	"dompetin-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "-" {
			return ""
		}
		return name
	})
}

func getErrorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return MsgRequired
	case "email":
		return MsgInvalidEmail
	case "min":
		return MsgMinValue + e.Param()
	case "max":
		return MsgMaxValue + e.Param()
	default:
		return MsgInvalidValue
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errs := c.Errors
		if len(errs) > 0 {
			err := errs.Last().Err

			switch e := err.(type) {
			case validator.ValidationErrors:
				validationErrors := make(map[string]string)
				for _, ve := range e {
					field := ve.Field()
					validationErrors[field] = getErrorMsg(ve)
				}
				response.Respond(c, http.StatusBadRequest, false, MsgErrorValidation, validationErrors, nil)
				return

			default:
				response.Respond(c, http.StatusInternalServerError, false, err.Error(), nil, nil)
				return
			}
		}
	}
}
