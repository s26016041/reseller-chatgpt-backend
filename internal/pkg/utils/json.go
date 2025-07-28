package utils

import (
	"errors"
	"fmt"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Tag Annotation        Source	                     Example Explanation
// uri:"id"	             Path Parameter	             e.g., GET /users/123 → route /users/:id
// json:"id"	         JSON Body	                 e.g., body content: {"id": 123}
// form:"id"	         Query or Form Parameter	 e.g., URL: ?id=123 or form field: id=123 (POST form)
// validate:"required"`  Required
var validate = validator.New()

func BindAll(ctx *gin.Context, obj interface{}) error {
	structTags := ParseBindingTags(obj)

	if structTags.Json {
		if err := ctx.ShouldBindJSON(obj); err != nil && !errors.Is(err, io.EOF) {
			return err
		}
	}

	if structTags.Uri {
		if err := ctx.ShouldBindUri(obj); err != nil {
			return fmt.Errorf("bind uri error: %w", err)
		}
	}

	if structTags.Form {
		if err := ctx.ShouldBindQuery(obj); err != nil {
			return fmt.Errorf("bind query error: %w", err)
		}
	}

	err := validate.Struct(obj)

	return err
}

type StructTag struct {
	Form bool
	Json bool
	Uri  bool
}

func ParseBindingTags(obj interface{}) StructTag {
	output := StructTag{}

	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if HasTag(field, "form") {
			output.Form = true
		}
		if HasTag(field, "json") {
			output.Json = true
		}
		if HasTag(field, "uri") {
			output.Uri = true
		}
	}
	return output
}

func HasTag(field reflect.StructField, tagName string) bool {
	_, ok := field.Tag.Lookup(tagName)
	return ok
}
