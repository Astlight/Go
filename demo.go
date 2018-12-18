package main

import (
	"fmt"
	"os"
)

//func main() {
//	fmt.Println(os.Args) // >>> 列表
//	fmt.Println(len(os.Args)) // >>> 1
//	// [C:\Users\HM-Python\AppData\Local\Temp\___go_build_demo_go.exe]
//}

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
