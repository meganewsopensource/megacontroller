package megacontroller

import (
	"github.com/invopop/jsonschema"
	"net/http"
)

type ResponseJsonParser interface {
	JSON(code int, obj any)
}

type H map[string]any

type Base struct{}

func (u Base) IncorrectSchema(c ResponseJsonParser, entity interface{}) {
	c.JSON(http.StatusUnprocessableEntity, H{
		"status":  "error",
		"message": "incorrect schema",
		"schema":  jsonschema.Reflect(entity),
	})
}

func (u Base) EmptyQueryParameter(c ResponseJsonParser, queryParameter string) {
	c.JSON(http.StatusBadRequest, H{
		"status":    "error",
		"message":   "empty query parameter",
		"parameter": queryParameter,
	})
}

func (u Base) NotIntegerQueryParameter(c ResponseJsonParser, queryParameter string) {
	c.JSON(http.StatusBadRequest, H{
		"status":    "error",
		"message":   "not integer query parameter",
		"parameter": queryParameter,
	})
}

func (u Base) EmptyHeader(c ResponseJsonParser, header string) {
	c.JSON(http.StatusBadRequest, H{
		"status":  "error",
		"message": "empty header",
		"header":  header,
	})
}
