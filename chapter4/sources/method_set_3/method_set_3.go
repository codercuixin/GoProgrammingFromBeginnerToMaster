package main

import "io"
import utils "mastergo/chapter4/sources/method_set_utils"


func main() {
	utils.DumpMethodSet((*io.Writer)(nil))
	utils.DumpMethodSet((*io.Reader)(nil))
	utils.DumpMethodSet((*io.Closer)(nil))
	utils.DumpMethodSet((*io.ReadWriter)(nil))
	utils.DumpMethodSet((*io.ReadWriteCloser)(nil))
}
