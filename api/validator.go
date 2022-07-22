package api

import (
	util "github.com/baoduong1011/Project_Golang/utils"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency,ok := fieldLevel.Field().Interface().(string); ok {
		// check currency is supported 
		return util.IsSupportedCurrency(currency)
	}

	return false
}