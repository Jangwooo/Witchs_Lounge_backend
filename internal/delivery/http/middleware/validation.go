package middleware

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBody[T any]() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body T
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "잘못된 요청 형식입니다",
			})
		}

		validate := validator.New()
		if err := validate.Struct(body); err != nil {
			var errs validator.ValidationErrors
			if errors.As(err, &errs) {
				details := make(map[string]string)
				t := reflect.TypeOf(body)
				for _, e := range errs {
					field, _ := t.FieldByName(e.StructField())
					jsonName := field.Tag.Get("json")
					if jsonName == "" {
						jsonName = e.Field()
					} else if idx := strings.Index(jsonName, ","); idx != -1 {
						jsonName = jsonName[:idx]
					}
					details[jsonName] = e.Tag()
				}
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   "유효하지 않은 요청 필드입니다",
					"details": details,
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Locals("body", body)
		return c.Next()
	}
}
