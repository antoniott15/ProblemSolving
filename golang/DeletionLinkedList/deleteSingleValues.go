package main

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func removeDuplicates(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head == nil {
		return nil
	}
	var temp *SinglyLinkedListNode = head.next
	for temp != nil && head.data == temp.data {
		temp = temp.next
	}
	head.next = removeDuplicates(temp)

	return head
}
