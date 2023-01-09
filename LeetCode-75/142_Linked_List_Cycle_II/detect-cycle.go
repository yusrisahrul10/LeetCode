package main

import "fmt"

func main() {
	list := ListNode{1, &ListNode{2, nil}}
	Show(detectCycle(&list))
}

func Show(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	slow := head
	fast := head
	isCycle := false

	for fast != nil && fast.Next != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	for head != slow {
		head = head.Next
		slow = slow.Next
	}

	return head
}
