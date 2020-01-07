package main

import (
	"fmt"
	"os"
)

/**
defer
保证函数最后执行该行
*/
func main() {

	// /Users/marszhang/log

	// 1。在使用createFile得到一个文件对象后，使用defer来调用关闭文件closeFIle，
	// 2。这个方法在main中最后执行，也就是在writeFile之后

	file := createFile("/Users/marszhang/log/go_defer_test.txt")
	defer closeFile(file)

	writeFile(file)
	fmt.Println("do other thing")

}

func createFile(p string) *os.File {

	fmt.Println("creating")
	f, err := os.Create(p) // 创建文件对象
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data联通")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
