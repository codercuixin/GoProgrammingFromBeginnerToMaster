package main
import utils "mastergo/chapter4/sources/method_set_utils"
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
	utils.DumpMethodSet(&t)
	utils.DumpMethodSet(&pt)
	utils.DumpMethodSet((*Interface)(nil))
}
