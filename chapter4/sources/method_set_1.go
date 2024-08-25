package main
import "fmt"
type Interface interface {
	M1()
	M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func main() {
	var t T
	var pt *T
	var i Interface
	fmt.Println(i)

	i = t
	i = pt
	
}
