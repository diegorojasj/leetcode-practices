package main

import (
	"sort"
	"reflect"
	"testing"
	"time"
)

func removeDuplicates(nums []int) int {
    encountered := map[int]bool{}
    unique := []int{}
    for i,v := range nums {
        if !encountered[v] {
            encountered[v] = true
            unique = append(unique, v)
        } else {
            nums[i] = 999999
        }
    }
    sort.Ints(nums)
    return len(unique)
}

func TestRemoveDuplicates(t *testing.T) {
    cases := []struct {
        nums     []int
        expected []int
    }{
        {[]int{1, 1, 2}, []int{1, 2}},
        {[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
    }
    for i, c := range cases {
		start := time.Now()
        result := removeDuplicates(c.nums)
        elapsed := time.Since(start)
		ns := elapsed.Nanoseconds()
        if result != len(c.expected) || !reflect.DeepEqual(c.nums[:result], c.expected) {
            t.Errorf("Case %d: removeDuplicates() = %v, nums[:k] = %v, expected %v in %dns", i+1, result, c.nums[:result], c.expected, ns)
        } else {
			t.Logf("Case %d: Passed in %dns", i+1, ns)
		}
    }
}
