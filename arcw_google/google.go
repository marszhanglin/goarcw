package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"strconv"
)

func main() {
	scanFile("./arcw_google/google.go")

}

//c:=3+4i  为复数
//fmt.Println(cmplx.Abs(c)) // 3+4i的绝对值
// 欧拉公式 E的i派次方加1 得 0    (0+1.2246467991473515e-16i)
func euler() {

	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)         // 注意浮点数是不准的
	fmt.Printf("%.10f\n", cmplx.Pow(math.E, 1i*math.Pi)+1) // 注意浮点数是不准的

	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.10f\n", cmplx.Exp(1i*math.Pi)+1) // 注意浮点数是不准的
}

//  求直角三角形斜边长 3平方加4平方
func Hypotenuse() {
	const a, b = 3, 4
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println("直角三角形斜边为：", c)
}

// int 转 二进制
//convertToBin(17899999)
func convertToBin(n int) string {
	result := ""
	data := n
	for ; n > 0; n /= 2 { // 对2取模，得到最低位，再除以2，循环。。
		lsb := n % 2 //对2取模
		result = strconv.Itoa(lsb) + result

	}
	fmt.Println(data, "转二进制为：", result)
	return result
}

// 扫描文件
func scanFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
