package service

import (
	"errors"

	"qwer/utils"
)

func IsStringEven(str string) (*bool, error) {
	if len(str) == 0 {
		return nil, errors.New("lenght of the string is 0")
	}
	if len(str)%2 != 0 {
		return utils.GetPointer(false), nil
	}
	return utils.GetPointer(true), nil
}