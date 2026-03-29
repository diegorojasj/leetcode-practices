# 2. Add Two Numbers

**Difficulty:** Medium

## Problem

You are given two **non-empty linked lists** representing two non-negative integers. The digits are stored in **reverse order**, and each node contains a single digit. Add the two numbers and return the sum as a linked list.

## Constraints

- The two numbers do **not** contain leading zeros, except for the number `0` itself.
- The number of nodes in each list is in the range `[1, 100]`.
- `0 <= Node.val <= 9`

## Examples

```
Input:  l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807
```

```
Input:  l1 = [0], l2 = [0]
Output: [0]
```

```
Input:  l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```