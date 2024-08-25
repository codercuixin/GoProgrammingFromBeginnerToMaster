package main
import "fmt"
type IntSliceFunctor interface{
	Fmap(fn func(int)int) IntSliceFunctor
}

type IntSliceFunctorImpl struct{
	ints []int 
}

func(isf IntSliceFunctorImpl) Fmap(fn func(int)int) IntSliceFunctor{
	newInts := make([]int, len(isf.ints))
	for index, value := range isf.ints {
		newValue := fn(value)
		//不能用 append 最后一个会触发扩容
		newInts[index] = newValue
	}
	return IntSliceFunctorImpl{ints: newInts}
}

func main(){
	var trippleFunc = func(x int) int {
		return x *3
	}
	var isf IntSliceFunctorImpl= IntSliceFunctorImpl{ints: []int{1, 2, 3, 4}}
	var trippeIsF = isf.Fmap(trippleFunc)
	fmt.Println(trippeIsF)
}