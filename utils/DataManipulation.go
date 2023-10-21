package utils

import (
	"fmt"

	"github.com/stoewer/go-strcase"
)

func ToCamelCase(str *string) {
	camelCase := strcase.LowerCamelCase(*str)
	fmt.Println(camelCase)
}
