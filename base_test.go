package megacontroller

import (
	"github.com/invopop/jsonschema"
	"net/http"
	"reflect"
	"testing"
)

type responseJsonParserMock struct {
	code int
	obj  any
}

func (r *responseJsonParserMock) JSON(code int, obj any) {
	r.code = code
	r.obj = obj
}

type myEntity struct {
	Property string `json:"property"`
}

func TestBase_IncorrectSchema(t *testing.T) {
	type args struct {
		c            *responseJsonParserMock
		entity       interface{}
		expectedCode int
		expectedObj  any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				c:            &responseJsonParserMock{},
				entity:       &myEntity{},
				expectedCode: http.StatusUnprocessableEntity,
				expectedObj: H{
					"status":  "error",
					"message": "incorrect schema",
					"schema":  jsonschema.Reflect(&myEntity{}),
				},
			},
		},
		{
			name: "Nil entity",
			args: args{
				c:            &responseJsonParserMock{},
				entity:       nil,
				expectedCode: http.StatusUnprocessableEntity,
				expectedObj: H{
					"status":  "error",
					"message": "incorrect schema",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Base{}
			u.IncorrectSchema(tt.args.c, tt.args.entity)
			if tt.args.expectedCode != tt.args.c.code {
				t.Errorf("IncorrectSchema() failed, expected code: %d, but got: %d", tt.args.expectedCode, tt.args.c.code)
			}
			if !reflect.DeepEqual(tt.args.expectedObj, tt.args.c.obj) {
				t.Errorf("IncorrectSchema() failed, expected obj: %v, but got: %v", tt.args.expectedObj, tt.args.c.obj)
			}
		})
	}
}

func TestBase_EmptyQueryParameter(t *testing.T) {
	const ValidQueryParameterName = "ValidQueryParameterName"
	type args struct {
		c              *responseJsonParserMock
		queryParameter string
		expectedCode   int
		expectedObj    any
	}
	tt := struct {
		name string
		args args
	}{
		name: "Success",
		args: args{
			c:              &responseJsonParserMock{},
			queryParameter: ValidQueryParameterName,
			expectedCode:   http.StatusBadRequest,
			expectedObj: H{
				"status":    "error",
				"message":   "empty query parameter",
				"parameter": ValidQueryParameterName,
			},
		},
	}

	t.Run(tt.name, func(t *testing.T) {
		u := &Base{}
		u.EmptyQueryParameter(tt.args.c, tt.args.queryParameter)
		if tt.args.expectedCode != tt.args.c.code {
			t.Errorf("EmptyQueryParameter() failed, expected code: %d, but got: %d", tt.args.expectedCode, tt.args.c.code)
		}
		if !reflect.DeepEqual(tt.args.expectedObj, tt.args.c.obj) {
			t.Errorf("EmptyQueryParameter() failed, expected obj: %v, but got: %v", tt.args.expectedObj, tt.args.c.obj)
		}
	})
}

func TestBase_NotIntegerQueryParameter(t *testing.T) {
	const ValidQueryParameterName = "ValidQueryParameterName"
	type args struct {
		c              *responseJsonParserMock
		queryParameter string
		expectedCode   int
		expectedObj    any
	}
	tt := struct {
		name string
		args args
	}{
		name: "Success",
		args: args{
			c:              &responseJsonParserMock{},
			queryParameter: ValidQueryParameterName,
			expectedCode:   http.StatusBadRequest,
			expectedObj: H{
				"status":    "error",
				"message":   "not integer query parameter",
				"parameter": ValidQueryParameterName,
			},
		},
	}

	t.Run(tt.name, func(t *testing.T) {
		u := &Base{}
		u.NotIntegerQueryParameter(tt.args.c, tt.args.queryParameter)
		if tt.args.expectedCode != tt.args.c.code {
			t.Errorf("EmptyQueryParameter() failed, expected code: %d, but got: %d", tt.args.expectedCode, tt.args.c.code)
		}
		if !reflect.DeepEqual(tt.args.expectedObj, tt.args.c.obj) {
			t.Errorf("EmptyQueryParameter() failed, expected obj: %v, but got: %v", tt.args.expectedObj, tt.args.c.obj)
		}
	})
}

func TestBase_EmptyHeader(t *testing.T) {
	const ValidHeaderName = "ValidQueryParameterName"
	type args struct {
		c              *responseJsonParserMock
		queryParameter string
		expectedCode   int
		expectedObj    any
	}
	tt := struct {
		name string
		args args
	}{
		name: "Success",
		args: args{
			c:              &responseJsonParserMock{},
			queryParameter: ValidHeaderName,
			expectedCode:   http.StatusBadRequest,
			expectedObj: H{
				"status":  "error",
				"message": "empty header",
				"header":  ValidHeaderName,
			},
		},
	}

	t.Run(tt.name, func(t *testing.T) {
		u := &Base{}
		u.EmptyHeader(tt.args.c, tt.args.queryParameter)
		if tt.args.expectedCode != tt.args.c.code {
			t.Errorf("EmptyQueryParameter() failed, expected code: %d, but got: %d", tt.args.expectedCode, tt.args.c.code)
		}
		if !reflect.DeepEqual(tt.args.expectedObj, tt.args.c.obj) {
			t.Errorf("EmptyQueryParameter() failed, expected obj: %v, but got: %v", tt.args.expectedObj, tt.args.c.obj)
		}
	})
}
