package main
import utils "mastergo/chapter4/sources/method_set_utils"

type Interface interface {
	M1()
	M2()
}

type T struct {
	Interface
}

func (T) M3() {}

func main() {
	utils.DumpMethodSet((*Interface)(nil))
	var t T = T{
	}
	var pt *T
	utils.DumpMethodSet(&t)
	utils.DumpMethodSet(&pt)

	//panic: runtime error: invalid memory address or nil pointer dereference
	t.M1();
	t.M2();
}
