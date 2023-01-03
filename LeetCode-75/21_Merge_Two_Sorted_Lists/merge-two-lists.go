package main

import "fmt"

func main() {
	// var head *ListNode
	// for i := 1; i <= 5; i++ {
	// 	head = insert(head, i)
	// }
	list1 := ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	// Show(mergeTwoLists(&list1, &list2))
	Show(mergeTwoLists(&list1, &list2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Show(l *ListNode) {
	p := l
	for p != nil {
		fmt.Printf("-> %v ", p.Val)
		p = p.Next
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1 := list1
	p2 := list2
	var newList *ListNode
	isNextNil := false

	if p2 == nil {
		return p1
	}

	if p1 == nil {
		return p2
	}

	for p2 != nil {
		for p1 != nil {
			if p1.Val < p2.Val || isNextNil {
				newList = insert(newList, p1.Val)
			} else {
				newList = insert(newList, p2.Val)
				if p2.Next == nil {
					isNextNil = true
					newList = insert(newList, p1.Val)
				} else {
					break
				}
			}
			p1 = p1.Next
		}
		if p1 == nil {
			break
		}
		p2 = p2.Next
	}
	for p2 != nil {
		if !isNextNil {
			newList = insert(newList, p2.Val)
		}
		p2 = p2.Next
	}

	return newList
}

func insert(root *ListNode, elem int) *ListNode {
	temp := ListNode{Val: elem, Next: nil}
	if root == nil {
		root = &temp
		return root
	}
	curr := root
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = &temp
	return root
}
