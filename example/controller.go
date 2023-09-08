package example

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/meganewsopensource/megacontroller"
	"strconv"
)

// Theses examples show how to use the errors from the megacontroller.Base
type exampleControler struct {
	megacontroller.Base
}

func (c exampleControler) HeaderExample(ctx *gin.Context) {
	const headerKey = "header"
	header := ctx.GetHeader(headerKey)
	if header == "" {
		c.EmptyHeader(ctx, headerKey)
		return
	}
}

func (c exampleControler) QueryExample(ctx *gin.Context) {
	const queryKey = "query"
	query := ctx.Query(queryKey)
	if query == "" {
		c.EmptyQueryParameter(ctx, queryKey)
		return
	}
}

func (c exampleControler) NotIntegerExample(ctx *gin.Context) {
	const integerQueryKey = "query"
	integerQueryStr := ctx.Query(integerQueryKey)
	_, err := strconv.Atoi(integerQueryStr)
	if err != nil {
		c.NotIntegerQueryParameter(ctx, integerQueryKey)
	}
}

func (c exampleControler) IncorrectSchemaExample(ctx *gin.Context) {
	user := &UserInput{}
	err := ctx.ShouldBindBodyWith(user, binding.JSON)
	if err != nil {
		c.IncorrectSchema(ctx, UserInput{})
	}
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
