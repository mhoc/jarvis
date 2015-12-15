package util

import (
	"jarvis/log"
	"math/rand"
)

var (
	letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func Check(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}

// Returns true if there is an error. Weird, right?
func CheckWarn(e error) bool {
	if e == nil {
		return false
	} else {
		log.Warn(e.Error())
		return true
	}
}

// Checks whether a map has the given elements as keys
// Returns false if it is missing any of the elements
func MapHasElements(m map[string]interface{}, elements ...string) bool {
	for _, element := range elements {
		if _, in := m[element]; !in {
			return false
		}
	}
	return true
}

func NewId() string {
	return NewIdn(10)
}

func NewIdn(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
