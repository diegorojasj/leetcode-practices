package main

import (
	"reflect"
	"testing"
	"time"
)

func isValid(s string) bool {
    stack := []rune{}
    dict := map[rune]rune {
        '(': ')',
        '{': '}',
        '[': ']',
    }
    for _, v := range s {
        elem, ok := dict[v]
        if ok {
            stack = append(stack, elem)
        } else {
            stackLen := len(stack)
            if stackLen == 0 {
                return false
            }
            pop := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if v != pop {
                return false
            }
        }
    }
    if len(s) > 0 && len(stack) == 0 {
        return true
    }
    return false
}

func TestIsValid(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"((", false},
		{"))", false},
	}

	for i, c := range cases {
		start := time.Now()
		result := isValid(c.s)
		elapsed := time.Since(start)
		ns := elapsed.Nanoseconds()
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Case %d: isValid(%s) = %v; expected %v in %dns", i+1, c.s, result, c.expected, ns)
		} else {
			t.Logf("Case %d: Passed in %dns", i+1, ns)
		}
	}
}
