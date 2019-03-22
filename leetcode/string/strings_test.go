package string

import (
	"fmt"
	"testing"
)

func TestShiftLetters(t *testing.T) {
	S := "abc"
	shifts := []int{3, 5, 9}
	res := shiftingLetters(S, shifts)
	fmt.Println(res)
}
