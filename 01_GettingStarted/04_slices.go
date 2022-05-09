package main

import "fmt"

type MyBoxItem struct {
	Name string
}

type MyBox struct {
	Items []MyBoxItem
}

func (box *MyBox) AddItem(item MyBoxItem) []MyBoxItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func (box *MyBox) AddItems(items ...MyBoxItem) []MyBoxItem {

	for _, item := range items {
		box.Items = append(box.Items, item)
	}
	return box.Items
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {

	/************** Slice of integer *************/
	// declare a slice
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	/************** Slice of structure *************/
	item1 := MyBoxItem{Name: "Item 1"}
	item2 := MyBoxItem{Name: "Item 2"}
	item3 := MyBoxItem{Name: "Item 3"}

	items := []MyBoxItem{}
	box := MyBox{items}

	box.AddItem(item1)
	box.AddItems(item2, item3)
	fmt.Println(box)
	fmt.Println(len(box.Items))
}
