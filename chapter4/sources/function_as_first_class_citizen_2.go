package main

import (
	"fmt"
	"net/http"
)

type MyFunc func(w http.ResponseWriter, r *http.Request)

func (f MyFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}

func main() {
	// http.ListenAndServe(":8080", MyFunc(greeting))
	//http.HandlerFunc(greeting) 也是一样的类型转换，将 greeting 转换成 MyFunc
	// http.ListenAndServe(":8080", http.HandlerFunc(greeting))
	intconvert()
}

func intconvert(){
	type MyInt int 
	var x int = 5
	y := MyInt(x)
	fmt.Println(y)
}
