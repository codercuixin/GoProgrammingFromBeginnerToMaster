package main
import utils "mastergo/chapter4/sources/method_set_utils"

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type Interface interface {
	M1()
	M2()
}

type T1 = T
type Interface1 = Interface

func main() {
	var t T
	var pt *T
	var t1 T1
	var pt1 *T1


	utils.DumpMethodSet(&t)
	utils.DumpMethodSet(&t1)

	utils.DumpMethodSet(&pt)
	utils.DumpMethodSet(&pt1)


	utils.DumpMethodSet((*Interface)(nil))
	utils.DumpMethodSet((*Interface1)(nil))
}
