package main

import (
	"reflect"
	"testing"
	"time"
)

type ListNode struct {
    Val  int
    Next *ListNode
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    current := dummy
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            current.Next = list1
            list1 = list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
    }
    if list1 != nil {
        current.Next = list1
    } else {
        current.Next = list2
    }
    return dummy.Next
}

func TestMergeTwoLists(t *testing.T) {
    cases := []struct {
        list1    *ListNode
        list2    *ListNode
        expected *ListNode
    }{
        {
            list1:    &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}},
            list2:    &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
            expected: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}},
        },
        {
            list1:    nil,
            list2:    nil,
            expected: nil,
        },
        {
			list1:    nil,
            list2:    &ListNode{Val: 0},
            expected: &ListNode{Val: 0},
        },
    }

	for i, c := range cases {
		start := time.Now()
		result := mergeTwoLists(c.list1, c.list2)
		elapsed := time.Since(start)
		ns := elapsed.Nanoseconds()
		if !reflect.DeepEqual(result, c.expected) {
			t.Errorf("Case %d: mergeTwoLists(%v, %v) = %v; expected %v in %dns", i+1, c.list1, c.list2, result, c.expected, ns)
		} else {
			t.Logf("Case %d: Passed in %dns", i+1, ns)
		}
	}
}
