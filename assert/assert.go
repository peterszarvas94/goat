package assert

import "fmt"

func Equal(expected, actual interface{}) {
	if expected != actual {
		panic(fmt.Sprintf("Expected %v, got %v", expected, actual))
	}
}

func True(condition bool) {
	if !condition {
		panic("Expected condition to be true")
	}
}

func False(condition bool) {
	if condition {
		panic("Expected condition to be false")
	}
}

func Nil(value interface{}) {
	if value != nil {
		panic(fmt.Sprintf("Expected nil, got %v", value))
	}
}

func NotNil(value interface{}) {
	if value == nil {
		panic("Expected value not to be nil")
	}
}

func Len(length int, str string) {
	if length != len(str) {
		panic(fmt.Sprintf("Expected %s to be length of %d, got %d", str, length, len(str)))
	}
}
