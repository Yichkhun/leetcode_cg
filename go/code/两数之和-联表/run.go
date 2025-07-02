package main

import "fmt"

//两数之和 leet code题解
//https://leetcode.cn/problems/add-two-numbers/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1->2>-3; 321
2->3; 32

321+32 = 353

题解的核心思路：
从左右指针遍历，因为起始的指针对应的就是求和结果的尾数，要考虑求和的过程中可能会进位。

1+2;
2+3;
3+0;
353
*/

func twoNumberSum(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode
	tail := new(ListNode)

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 3}}

	res := twoNumberSum(l1, l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
