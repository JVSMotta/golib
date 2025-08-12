package golib

import ("fmt"
	"runtime")

func version(){
	fmt.Println(runtime.Version())
}
