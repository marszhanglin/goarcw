package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/**
 */
func main() {
	// 1.将字符串写入文件
	//write1()

	//write2()

	//write3()

	//write4()

}

func write4() {
	f, err := os.Create("/Users/marszhang/log/go_file_test4.txt")
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Println("wrote %d bytes\n", n4)
	w.Flush()
}

func write3() {
	f, err := os.Create("/Users/marszhang/log/go_file_test3.txt")
	check(err)
	defer f.Close()
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Println("Write %d bytes \n", n3)
	f.Sync() // 调用Sync方法来将缓冲区数据写入磁盘
}

func write2() {
	f, err := os.Create("/Users/marszhang/log/go_file_test2.txt")
	check(err)
	defer f.Close()
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Println(" write %d bytes\n", n2)
}

func write1() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/Users/marszhang/log/go_file_test1.txt", d1, 0664)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
