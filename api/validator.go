package api

import "github.com/go-playground/validator/v10"
import "github.com/techschool/simplebank/util"

var validCurrecy validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
