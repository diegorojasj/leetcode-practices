package main

import (
	"reflect"
	"testing"
	"time"
)

import (
    "strings"
)
func longestCommonPrefix(strs []string) string {
    result := ""
    firstStrLength := len(strs[0])
    splitStrs := make([][]string, len(strs))
    for inx, str := range strs {
        splitStrs[inx] = strings.Split(str,"")
    }
    i:=0
    for i=0; i<firstStrLength; i++ {
        letter := ""
        br := false
        for inx, str := range splitStrs {
            if i >= len(str) {
                br = true
                break
            }
            if inx == 0 && str[i] != "" {
                letter = str[i]
            } else if str[i] == "" {
                br = true
                break
            }
            if letter != str[i] {
                br = true
                break
            }
        }
        if br {
            break
        }
        result+=splitStrs[0][i]
    }
    return result
}

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		strs []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		// extra cases
		{[]string{"ab", "a"}, "a"},
		{[]string{"alone"}, "alone"},
		{[]string{"same", "same", "same"}, "same"},
		{[]string{"xyz", "abc", "mno"}, ""},
		{[]string{"prefix", "prefixed", "prefixes"}, "prefix"},
		{[]string{"a", "a", "a"}, "a"},
		{[]string{"a", "b", "c"}, ""},
		{[]string{"123abc", "123def", "123"}, "123"},
		{[]string{"abcdefghij", "abcdefghiz", "abcdefghix"}, "abcdefghi"},
		{[]string{"flower", "first", "fast"}, "f"},
	}

	for i, c := range cases {
		start := time.Now()
		result := longestCommonPrefix(c.strs)
		elapsed := time.Since(start)
		ns := elapsed.Nanoseconds()
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Case %d: longestCommonPrefix(%v) = %v; expected %v in %dns", i+1, c.strs, result, c.expected, ns)
		} else {
			t.Logf("Case %d: Passed in %dns", i+1, ns)
		}
	}
}