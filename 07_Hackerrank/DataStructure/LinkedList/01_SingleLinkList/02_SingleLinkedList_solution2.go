/* Reference:
	* Example 1: Print all data in singly linked list.Refer: [here](https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list/problem)
	* Example 2: Insert at head. Refer: [here](https://www.hackerrank.com/challenges/insert-a-node-at-the-head-of-a-linked-list/problem)
	* Example 3: Insert at tail. Refer: [here](https://www.hackerrank.com/challenges/insert-a-node-at-the-tail-of-a-linked-list/problem)
	* Example 4: Insert at position. Refer: [here](https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem)
	* Example 5: Delete node at head. Refer: [here]()
	* Example 6: Delete node at tail.
	* Example 7: Delete node at middle.
	* Example 8: Delete node at position.
	* Example 9: Delete node where data = input key.
	* Example 10: Print reverse value a Singly Linked List by recursive expression
	* Example 11: Print reverse value a Singly Linked List by dynamic array

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
		fmt.Printf("%d -> ", head.data)
		head = head.next
	}
	fmt.Println()
}

func (singlyLinkedList *SinglyLinkedList) printLength() {
	var length int32 = 0
	head := singlyLinkedList.head
	for head != nil {
		length++
		head = head.next
	}
	fmt.Println("Length of Singly Linked List: ", length)
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

// Example 5: Complete the deleteNodeAtHead function below.
func (singlyLinkedList *SinglyLinkedList) deleteNodeAtHead() {

	// temp := singlyLinkedList.head

	singlyLinkedList.head = singlyLinkedList.head.next

	// Step 3: We can free memory allocation for temp by assiging nil
	//temp = nil
}

// Example 6: Complete the deleteNodeAtEnd function below.
func (singlyLinkedList *SinglyLinkedList) deleteNodeAtEnd() {

}

// Example 7: Complete the deleteNodeAtMiddle function below.
func (singlyLinkedList *SinglyLinkedList) deleteNodeAtMiddle(position int32) {
	// Demo input singly linked list
	/*	Head									  Tail
		---------------		---------------		---------------
		| 16 | Next |	->	| 13 | Next |	-> 	| 7 | Next | -> nil
		---------------		---------------		---------------
			NODE1				NODE2				NODE3
	*/
	temp := singlyLinkedList.head
	prev := singlyLinkedList.head
	for i := int32(0); i < position; i++ {
		if i == 0 && position == 1 {
			singlyLinkedList.head = singlyLinkedList.head.next
		} else {
			if i == position-1 && temp != nil {
				prev.next = temp.next
			} else {
				prev = temp
				if prev == nil { // position was greater than number of nodes in the list
					break
				}
				temp = temp.next
			}
		}
	}

}

// Example 10: Complete the reversePrintByRecursive function below.
func reversePrintByRecursive(llist *SinglyLinkedListNode) {

	if llist == nil {
		return
	} else {
		reversePrintByRecursive(llist.next)
		fmt.Printf("%d\n", llist.data)
	}
}

// Example 11: Complete the reversePrintByStack function below.
func reversePrintByStack(llist *SinglyLinkedListNode) {

	curr := llist
	// Initialize a resizable array
	stack := []int32{}
	if llist == nil {
		return
	} else {
		// Insert data to stack
		for curr != nil {
			stack = append(stack, curr.data)
			curr = curr.next
		}

		// Reorder number to print value
		for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
			stack[i], stack[j] = stack[j], stack[i]
		}

	}

	// Display reverse data in Linked List
	for i := range stack {
		fmt.Printf("%d\n", stack[i])
	}

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

	fmt.Println("Example 2: Insert at head 2 node with data = 1 and data = 2")
	llist.insertNodeAtHead(1)
	llist.insertNodeAtHead(2)
	printLinkedList(llist.head)

	fmt.Println("Example 3: Insert at tail 2 node with data = 19  and data = 20")
	llist.insertNodeAtTail(19)
	llist.insertNodeAtTail(20)
	printLinkedList(llist.head)

	fmt.Println("Example 4: Insert at position 2 by a new node with data = 5")
	llist_after := insertNodeAtPosition(llist.head, 5, 2)
	printLinkedList(llist_after)

	fmt.Println("Example 4: Delete node at head")
	llist.deleteNodeAtHead()
	llist.printLength()
	printLinkedList(llist.head)

	fmt.Println("Example 6: Delete node at position")
	llist.deleteNodeAtMiddle(4)
	llist.printLength()
	printLinkedList(llist.head)

	fmt.Println("Example 10: Print reverse value a Singly Linked List by recursive expression")
	reversePrintByRecursive(llist.head)

	fmt.Println("Example 11: Print reverse value a Singly Linked List by dynamic array")
	reversePrintByStack(llist.head)
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
