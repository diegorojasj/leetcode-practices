package main

import (
	"reflect"
	"testing"
	"time"
)

func twoSum(nums []int, target int) []int {
	mapping := map[int]int{}
	for i, num := range nums {
		rest := target - num
		if ind, found := mapping[rest]; found {
			return []int{ind, i}
		}
		mapping[num] = i
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{3, 2, 3}, 6, []int{0, 2}},
	}

	for i, c := range cases {
		start := time.Now()
		result := twoSum(c.nums, c.target)
		elapsed := time.Since(start)
		ns := elapsed.Nanoseconds()
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Case %d: twoSum(%v, %d) = %v; expected %v in %dns", i+1, c.nums, c.target, result, c.expected, ns)
		} else {
			t.Logf("Case %d: Passed in %dns", i+1, ns)
		}
	}
}
