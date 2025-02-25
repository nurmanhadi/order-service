package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Success(ctx *fiber.Ctx, statusCode int, data any) error {
	if data == nil {
		return ctx.Status(statusCode).JSON(fiber.Map{
			"links": fiber.Map{
				"self": ctx.OriginalURL(),
			},
		})
	}
	return ctx.Status(statusCode).JSON(fiber.Map{
		"data": data,
		"links": fiber.Map{
			"self": ctx.OriginalURL(),
		},
	})
}
func Error(ctx *fiber.Ctx, err error) error {
	if validatorErr, ok := err.(validator.ValidationErrors); ok {
		var values []string
		for _, fieldErr := range validatorErr {
			value := fmt.Sprintf("field %s is %s %s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param())
			values = append(values, value)
		}
		str := strings.Join(values, ", ")
		return ErrorR(ctx, 400, str)
	}
	return ErrorR(ctx, 500, err.Error())
}
func ErrorR(ctx *fiber.Ctx, statusCode int, err string) error {
	return ctx.Status(statusCode).JSON(fiber.Map{
		"error": err,
		"links": fiber.Map{
			"self": ctx.OriginalURL(),
		},
	})
}
