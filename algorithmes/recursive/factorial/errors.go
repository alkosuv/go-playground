package factorial

import "errors"

var (
	ErrNegativNumber = errors.New("число должно быть неотрицательным")
	ErrOverflow      = errors.New("переполнение: факториал слишком велик для int64")
)
