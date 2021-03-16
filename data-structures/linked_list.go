package data_structures

import "fmt"

type LinkedList struct {
	head *ListNode
}

type ListNode struct {
	value int
	next  *ListNode
}

func (l *LinkedList) print() {
	n := l.head
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

func (l *LinkedList) add(node ListNode) {
	n := l.head

	if n == nil {
		l.head = &node
	} else {
		for n.next != nil {
			n = n.next
		}
		n.next = &node
	}
}

func (l *LinkedList) tail() *ListNode {
	n := l.head
	if n == nil {
		return nil
	} else {

		for n.next != nil {
			n = n.next
		}
		return n
	}
}

func (l *LinkedList) reverse() {
	var prev *ListNode

	n := l.head

	for n != nil {
		curr := n
		n = n.next

		curr.next = prev
		l.head = curr
		prev = curr
	}
}
