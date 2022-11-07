/* Reference:
	* Example 1: Print all data in singly linked list.Refer: [here](https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list/problem)
	* Example 2: Insert at head. Refer: [here](https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list/problem)
	* Example 3: Insert at tail. Refer: [here](https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list/problem)
	* Example 4: Insert at position. Refer: [here](https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem)
	* Example 5:

Sample Input:
2
16
13

Sample output:
16
13
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {

	// Step 1: Create a new node
	new_node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	// Step 2: Check node at head is nil? If nil insert new node to head : if not nil insert new_node to tail
	/*
		Case 1: Head = nil
		Head
		---------------
		| Data | Next |
		---------------
			new_node

		Case 2: Head != nil
		---------------		---------------		---------------
		| Data | Next |	->	| Data | Next |	-> 	| Data | Next |
		---------------		---------------		---------------
			NODE1				NODE2				new_node
	*/
	if singlyLinkedList.head == nil {
		singlyLinkedList.head = new_node
	} else {
		singlyLinkedList.tail.next = new_node
	}

	// Step 3: Assign tail equal a singly Linked list
	singlyLinkedList.tail = new_node
}

// Complete the printLinkedList function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func printLinkedList(head *SinglyLinkedListNode) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

// Complete the insertNodeAtTail function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func (singlyLinkedList *SinglyLinkedList) insertNodeAtTail(nodeData int32) {

	// Step 1: Create a new node
	new_node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}
	// Step 2: Assign new_node.next to a singly linked list
	/*		Head				  Tail
			---------------		---------------		---------------
			| Data | Next |	->	| Data | Next |	-> 	| Data | Next | -> nil
			---------------		---------------		---------------
				NODE1				NODE2				new_node
	*/
	if singlyLinkedList.head == nil {
		singlyLinkedList.head = new_node
	} else {
		singlyLinkedList.tail.next = new_node
	}
	// Step 3: Move tail to new_node
	/*		Head				  					     Tail
			---------------		---------------		---------------
			| Data | Next |	->	| Data | Next |	-> 	| Data | Next | -> nil
			---------------		---------------		---------------
				NODE1				NODE2				new_node
	*/
	singlyLinkedList.tail = new_node
}

// Complete the insertNodeAtHead function below.
/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func (singlyLinkedList *SinglyLinkedList) insertNodeAtHead(nodeData int32) {

	// Step 1: Create a new node
	new_node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}
	// Step 2: Assign new_node.next to a singly linked list
	/*						Head				  Tail
	---------------		---------------		---------------
	| Data | Next |	->	| Data | Next |	-> 	| Data | Next | -> nil
	---------------		---------------		---------------
		new_node			NODE1				NODE2
	*/
	new_node.next = singlyLinkedList.head

	// Step 3: Move head to new_node
	/*	  Head				  					  Tail
	---------------		---------------		---------------
	| Data | Next |	->	| Data | Next |	-> 	| Data | Next | -> nil
	---------------		---------------		---------------
		new_node			NODE1				NODE2
	*/
	singlyLinkedList.head = new_node

}

// Complete the insertNodeAtPosition function below.
/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {

	// Demo input singly linked list
	/*	Head									  Tail
		---------------		---------------		---------------
		| 16 | Next |	->	| 13 | Next |	-> 	| 7 | Next | -> nil
		---------------		---------------		---------------
			NODE1				NODE2				NODE3
	*/

	// Step 1: Create a new node
	new_node := SinglyLinkedListNode{
		data: data,
		next: nil,
	}

	// Step 2: Move to node at position 2
	/*	Head									  Tail
		---------------		---------------		---------------
		| 16 | Next |	->	| 13 | Next |	-> 	| 7 | Next | -> nil
		---------------		---------------		---------------
			NODE1				NODE2				NODE3
	*/
	curr := llist
	var i int32 = 0
	for i < position-1 {
		i++
		curr = curr.next
	}

	new_node.next = curr.next
	curr.next = &new_node

	return llist
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist := SinglyLinkedList{}
	for i := 0; i < int(llistCount); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist.insertNodeIntoSinglyLinkedList(llistItem)
	}

	/// Print all data in single list
	fmt.Println("Example 1: Print all data in Single List")
	printLinkedList(llist.head)

	fmt.Println("Example 2: Insert at head")
	llist.insertNodeAtHead(1)
	llist.insertNodeAtHead(2)
	printLinkedList(llist.head)

	fmt.Println("Example 3: Insert at tail")
	llist.insertNodeAtTail(19)
	llist.insertNodeAtTail(20)
	printLinkedList(llist.head)

	fmt.Println("Example 4: Insert at position")
	llist_after := insertNodeAtPosition(llist.head, 5, 2)
	printLinkedList(llist_after)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
