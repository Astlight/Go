package main

//func main() {
//	fmt.Println(os.Args) // >>> 列表
//	fmt.Println(len(os.Args)) // >>> 1
//	// [C:\Users\HM-Python\AppData\Local\Temp\___go_build_demo_go.exe]
//}

//func main() {
//	s, sep := "", ""
//	for _, arg := range os.Args[1:] {
//		s += sep + arg
//		sep = " "
//	}
//	fmt.Println(s)
//}

type Retiever interface {
	Get(url string) string // interface中 都是函数
}

func download(r Retiever) string {
	return r.Get("astlight")
}

func main() {
}
