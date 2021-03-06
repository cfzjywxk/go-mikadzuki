package util

import (
	"fmt"
)

func AssertEQ(left, right interface{}) {
	if left != right {
		panic(fmt.Sprintf("%v(%T) != %v(%T)", left, left, right, right))
	}
}

func AssertNE(left, right interface{}) {
	if left == right {
		panic(fmt.Sprintf("%v(%T) == %v(%T)", left, left, right, right))
	}
}

func AssertNil(i interface{}) {
	if i != nil {
		panic(fmt.Sprintf("%v is not nil", i))
	}
}

func AssertNotNil(i interface{}) {
	if i == nil {
		panic(fmt.Sprintf("%v is nil", i))
	}
}
