package main

import "github.com/Pallinder/go-randomdata"

func GetFullName(gender int) string {
	return randomdata.FullName(gender)
}

func GetRandomRunes(length int) string {
	return randomdata.RandStringRunes(length)
}
