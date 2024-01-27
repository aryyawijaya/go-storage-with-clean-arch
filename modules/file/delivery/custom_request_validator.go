package filedelivery

import (
	"github.com/aryyawijaya/go-storage-with-clean-arch/entity"
	"github.com/go-playground/validator/v10"
)

var ValidAccess validator.Func = func(fl validator.FieldLevel) bool {
	if access, ok := fl.Field().Interface().(entity.Access); ok {
		return access.Valid()
	}

	return false
}
