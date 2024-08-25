package main 
import "fmt"
func main(){
	typeconvert()
}

func typeconvert(){
	var a int = 5 
	var b int32 = 6 
	fmt.Println(a+int(b))
}

