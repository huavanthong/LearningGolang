package main

func main() {
	functions := []func(){}
	functions = append(functions, doSomething1)
	functions = append(functions, doSomething2)
	functions = append(functions, doSomething3)

	for _, f := range functions {
		go f()
	}
}

func doSomething1() {
	// Thực hiện công việc nào đó
}

func doSomething2() {
	// Thực hiện công việc nào đó
}

func doSomething3() {
	// Thực hiện công việc nào đó
}
