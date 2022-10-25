package helpers

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"` //karena isi dari data tidak selalu sama jadi menggunakan type interface{}
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonresponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonresponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) { //untuk mengubah error menjadi tipe validasi
		errors = append(errors, e.Error())
	}
	return errors
}
