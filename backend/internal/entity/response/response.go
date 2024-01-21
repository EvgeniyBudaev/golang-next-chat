package response

import (
	"errors"
	errorDomain "github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/error"
	"github.com/EvgeniyBudaev/golang-next-chat/backend/internal/entity/success"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func WrapError(ctf *fiber.Ctx, err error, httpStatusCode int) error {
	var customError *errorDomain.CustomError
	if errors.As(err, &customError) {
		msg := errorDomain.ResponseError{
			StatusCode: customError.StatusCode,
			Success:    false,
			Message:    customError.Err.Error(),
		}
		return ctf.Status(customError.StatusCode).JSON(msg)
	}
	msg := errorDomain.ResponseError{
		StatusCode: httpStatusCode,
		Success:    false,
		Message:    err.Error(),
	}
	return ctf.Status(httpStatusCode).JSON(msg)
}

func WrapOk(ctx *fiber.Ctx, data interface{}) error {
	msg := success.Success{
		Data:       data,
		StatusCode: http.StatusOK,
		Success:    true,
	}
	return ctx.Status(fiber.StatusOK).JSON(msg)
}

func WrapCreated(ctf *fiber.Ctx, data interface{}) error {
	msg := success.Success{
		Data:       data,
		StatusCode: http.StatusCreated,
		Success:    true,
	}
	return ctf.Status(fiber.StatusCreated).JSON(msg)
}
