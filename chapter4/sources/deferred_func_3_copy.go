package main

import "fmt"

func main() {
	foo()
	fmt.Println("main exit normally")
}

func foo() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("recoverd from painc")
		}
	}()

	bar()

}
func bar(){
	fmt.Println("raise a panic")
	panic(-1)
}