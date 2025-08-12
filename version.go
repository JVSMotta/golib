package version 

import ("fmt"
	"runtime")

func version(){
	fmt.Println(runtime.Version())
}
