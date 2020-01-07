package main

import "fmt"

/**
make
*/
func main() {
	// 和数组不同的是，<<切片的长度是可变的>>，我们可以使用内置函数make来创建一个长度不为零的切片
	// 这里我们创建了一个长度为3，存储字符串的切片，切片元素默认为零值，对于字符串就是""
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s:", s)
	fmt.Println("s[2]:", s[2])
	fmt.Println("len(s):", len(s))

	s = append(s, "d")
	fmt.Println("apd:", s)
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 拷贝
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// 截取
	l1 := s[2:5]
	fmt.Println("s[2:5]:", l1)

	l2 := s[:5]
	fmt.Println("s[:5]:", l2)

	l3 := s[2:]
	fmt.Println("s[2:]:", l3)

	// 也是切片？
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 多维切片 和数组不同的是，<<切片的长度是可变的>>
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
