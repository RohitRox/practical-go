package data_structures

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("LinkedList", func(t *testing.T) {

		list := LinkedList{}
		node1 := ListNode{value: 1}
		node2 := ListNode{value: 4}
		node3 := ListNode{value: 8}
		node4 := ListNode{value: 16}

		list.add(node1)
		list.add(node2)
		list.add(node3)
		list.add(node4)

		fmt.Printf("Head: %d \n", list.head.value)
		fmt.Printf("Tail: %d \n", list.tail().value)

		fmt.Println("List print:")
		list.print()

		list.reverse()

		fmt.Println("List reversed print:")
		fmt.Printf("Head: %d \n", list.head.value)
		fmt.Printf("Tail: %d \n", list.tail().value)

		list.print()
	})
}
