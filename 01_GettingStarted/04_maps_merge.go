/*
Analyze Requirement:
	m1 := map[string]string{"id_1": "val_1"}
	m2 := map[string]string{"id_2": "val_2"}
	m3 := map[string]string{"id_1": "val_3"}

Result:
	result_map = [id: id_1 val: {val_1, val_3}, id: id_2 var: {val_2}} ]

Abnormal:
	m4 := map[string]string{"id_1": "val_1"} ????????

Reference:
	https://stackoverflow.com/questions/50324612/merge-maps-in-golang
*/
package main

import "fmt"

func main() {
	m1 := map[string]string{"id_1": "val_1"}
	m2 := map[string]string{"id_2": "val_2"}
	m3 := map[string]string{"id_1": "val_3"}
	m4 := map[string]string{"id_1": "val_3"}

	res := merge(m1, m2, m3, m4)
	fmt.Println(res)

	res2 := merge2(m1, m2, m3, m4)
	fmt.Println(res2)
}

/*
Function name: merge

Usage:
	+ to merge multiple maps[string]string to a final maps

Arguement: ...map[string]string: get multiple maps as argument

Return: map[string][]string: return a maps with
	+ map[string]: key is a string
	+ []string: value is a array
*/
func merge(ms ...map[string]string) map[string][]string {
	// init a empty object
	res := map[string][]string{}

	// loop over arguments
	for _, m := range ms {
		// for each res[k] append value if have multiple value
		for k, v := range m {
			res[k] = append(res[k], v)
		}
	}
	return res
}

/*
Function name: merge2

Usage:
	+ to merge multiple maps[string]string to a final maps
	+ improve merge() with feature to avoid duplicate value

Arguement: ...map[string]string
	+ get multiple maps as argument

Return: map[string][]string
	return a maps with
	+ map[string]: key is a string
	+ []string: value is a array
*/
func merge2(ms ...map[string]string) map[string][]string {
	res := map[string][]string{}
	for _, m := range ms {
	srcMap:
		for k, v := range m {
			// Check if (k,v) was added before:
			for _, v2 := range res[k] {
				if v == v2 {
					continue srcMap
				}
			}
			res[k] = append(res[k], v)
		}
	}
	return res
}
