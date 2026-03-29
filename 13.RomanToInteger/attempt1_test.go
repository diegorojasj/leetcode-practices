package main

import (
	"reflect"
	"testing"
	"time"
)

func romanToInt(s string) int {
	result:= 0
    romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

    for i:=len(s) -1; i>=0; i-- {
        val := romanMap[s[i]]
        if i < len(s)-1 && val < romanMap[s[i+1]] {
            result -= val
        } else {
            result += val
        }
    }

    return result
}

func TestRomanToInt(t *testing.T) {
	cases := []struct {
		s       string
		expected int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, c := range cases {
		start := time.Now()
		result := romanToInt(c.s)
		elapsed := time.Since(start)
		ns := elapsed.Nanoseconds()
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Case %d: romanToInt(%s) = %d; expected %d in %dns", i+1, c.s, result, c.expected, ns)
		} else {
			t.Logf("Case %d: Passed in %dns", i+1, ns)
		}
	}
}