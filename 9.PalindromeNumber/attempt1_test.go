package main

import (
	"strconv"
	"reflect"
	"testing"
	"time"
)

func isPalindrome(x int) bool {
    str := strconv.Itoa(x)
    rev := Reverse(str)
    return rev == str
}

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		num     int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for i, c := range cases {
		start := time.Now()
		result := isPalindrome(c.num)
		elapsed := time.Since(start)
		ns := elapsed.Nanoseconds()
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Case %d: isPalindrome(%d) = %v; expected %v in %dns", i+1, c.num, result, c.expected, ns)
		} else {
			t.Logf("Case %d: Passed in %dns", i+1, ns)
		}
	}
}