package main

import "fmt"

type MyInt *int

// func(myInt MyInt) hello(){
// 	fmt.Println(myInt)
// }
type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) {
	t.a = a
}

func main() {
	 var t = T{a: 10}
	 fmt.Println(t.Get())
	 t.Set(100)
	 fmt.Println(t.Get())

	 f1 := (*T).Set 
	 f1(&t, 1000)
	 f2 := (T).Get
	 fmt.Println(f2(t))
}


