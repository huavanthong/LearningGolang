package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

type Node struct {
	data int
	next *Node
}

func push(head_ref **Node, data int) {

	new_node := &Node{
		data: data,
		next: (*head_ref),
	}

	(*head_ref) = new_node
}

func insertAfterNode(prev_node *Node, new_data int) {

	if prev_node == nil {
		fmt.Println("the given previous node cannot be NULL")
		return
	}

	new_node := &Node{
		data: new_data,
		next: prev_node.next,
	}

	prev_node.next = new_node
}

func insertAtPostion(head **Node, data int, postion int) {

	if postion == 0 {
		fmt.Println("Plesae choose postion from 1 -> n")
		return
	}
	new_node := &Node{
		data: data,
		next: nil,
	}

	node_pos := *head

	var i int = 1
	for i < postion {
		i++
		node_pos = node_pos.next

		log.WithFields(log.Fields{
			"Position": i,
			"Data":     node_pos.data,
		}).Info("Loop over list until at position")
	}

	new_node.next = node_pos.next
	node_pos.next = new_node
}

func append(head_ref **Node, new_data int) {
	new_node := &Node{
		data: new_data,
		next: nil,
	}

	last := *head_ref

	if *head_ref == nil {
		*head_ref = new_node
		return
	}

	for last != nil {
		last = last.next
	}

	last.next = new_node

}

func printList(head *Node) {

	temp := head

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main() {

	var head *Node = nil
	append(&head, 6)
	push(&head, 4)
	push(&head, 1)
	printList(head)
	fmt.Println("#################")

	insertAfterNode(head, 2)
	insertAfterNode(head.next, 3)
	printList(head)

	fmt.Println("#################")
	push(&head, 0)
	printList(head)

	fmt.Println("#################")
	insertAtPostion(&head, 5, 5)
	printList(head)
}
