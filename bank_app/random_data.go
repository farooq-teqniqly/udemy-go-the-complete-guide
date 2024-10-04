package main

import (
	"errors"
	"github.com/Pallinder/go-randomdata"
)

func GetFullName(gender int) (string, error) {
	if gender != 0 && gender != 1 {
		return "", errors.New("gender must be 0 or 1")
	}

	return randomdata.FullName(gender), nil
}

func GetRandomRunes(length int) (string, error) {
	if length < 1 {
		return "", errors.New("length must be greater than 0")
	}

	return randomdata.RandStringRunes(length), nil
}
