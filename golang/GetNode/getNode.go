package main

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

func getNode(head *SinglyLinkedListNode, positionFromTail int32) int32 {
	var size, newSize int32 = 1, 1
	var temp *SinglyLinkedListNode = head
	for head != nil {
		size++
		head = head.next
	}

	if size == 1 {
		return temp.data
	}

	var newPos int32 = (size - positionFromTail)

	for temp != nil {
		newSize++
		if newSize == newPos {
			return temp.data
		}
		temp = temp.next
	}
	return 0
}
