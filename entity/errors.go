package entity

import "errors"

var (
	ErrUnique                  = errors.New("any field in your resource already existed")
	ErrNotFound                = errors.New("your requested resource is not found")
	ErrForbiddenAccess         = errors.New("forbidden to access resource")
	ErrNotSameLenSlice         = errors.New("length of 2 list is not equal")
	ErrEmptyPayload            = errors.New("your request payload should not be empty")
	ErrAnyEntityNotCreatedToDb = errors.New("any entity not created to db")
)
