package lib

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status           int          `json:"status"`                                           // http status
	Message          string       `json:"message"`                                          // response message
	ErrorDescription *string      `json:"error_description,omitempty" swaggerignore:"true"` // Error description
	ErrorData        *[]ErrorData `json:"error_data,omitempty" swaggerignore:"true"`        // Error fields
}

type ErrorData struct {
	// Namespace string      `json:"namespace" example:"User.Username"`                       // data namespace
	Name      string      `json:"name" example:"Username"`                                 // data name
	Path      string      `json:"path" example:"user.username"`                            // object property path
	Type      string      `json:"type,omitempty" example:"string"`                         // data type
	Value     interface{} `json:"value,omitempty" swaggertype:"string" example:"jane doe"` // value
	Validator string      `json:"validator" example:"required"`                            // validator type, see [more details](https://github.com/go-playground/validator#baked-in-validations)
	Criteria  interface{} `json:"criteria,omitempty" swaggertype:"number" example:"10"`    // criteria, example: if validator is gte (greater than) and criteria is 10, then it means a maximum of 10
	Message   string      `json:"message,omitempty" example:"invalid value"`               // Field message
}

// OK send http 200 response
func OK(c *fiber.Ctx, result ...interface{}) error {
	if len(result) == 0 {
		result = append(result, Response{
			Status:  200,
			Message: "success",
		})
	}

	return c.Status(200).JSON(result[0])
}
