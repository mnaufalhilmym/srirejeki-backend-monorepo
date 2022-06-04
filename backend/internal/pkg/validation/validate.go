package validation

import (
	"database/sql/driver"
	"reflect"
)

// type errorResponse struct {
// 	FailedField string `json:"failedField"`
// 	Tag         string `json:"tag"`
// 	Value       string `json:"value"`
// }

func ValidateStruct(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {

		val, err := valuer.Value()
		if err != nil {
			return err
		}
		return val
	}
	return nil
}
