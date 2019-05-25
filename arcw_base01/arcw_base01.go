// hello_2_grammer project hello_2_grammer.go
package main

//https://www.jianshu.com/p/bed39de53087
//本项目主要用来学习语法
import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println(math.Pi)
	a, b := param("张", "林")
	fmt.Println(a + b)
	fmt.Println(param("张", "林"))
	convert()
	default_zero()
	type_is()
	const_()
	for_()
	for_while()
	fmt.Println(if_sqrt(2), if_sqrt(-4))
	switch_()
	defer_()
	shuzhu_()
	slice_qp()
	range_()
	map_()
	func_test()
	fmt.Println(time.Now().UnixNano())

	fmt.Println(rand.Intn(1000))

	sha256test()

}

func sha256test() {
	h := sha256.New()
	h.Write([]byte("YzcmCZNvbXocrsz9dm8eadmin"))
	fmt.Printf("%x", h.Sum(nil))
}

//1-4.1-5.1-6.1-7 多个入参，多个返回值
func param(x, y string) (string, string) {
	return x, y
}

//1-12. 类型转换
func convert() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

//1-13.零值
func default_zero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

//1-14.类型推导
func type_is() {
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	b := "aa" // change me!
	fmt.Printf("b is of type %T\n", b)
}

//1-15.常量
const Pi = 3.14

func const_() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

//2-1 for 区别是少了()
func for_() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//2-3.for 是 Go 的 “while”
func for_while() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

//2-4.if
func if_sqrt(x float64) string {
	if x < 0 {
		return if_sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//2-7.if
func switch_() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("\nWhen's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}

//2-10. defer defer 语句会延迟函数的执行直到上层函数返回  也就是这个函数执行完毕后执行，或者说最后执行
func defer_() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

//3-2.结构体
//type Vertex struct {
//	X int
//	Y int
//}

//3-6.数组
func shuzhu_() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

//3-7.slice
func slice_() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

//3-7.slice_qp
func slice_qp() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])

	// 省略下标代表从 0 开始
	fmt.Println("s[:3] ==", s[:3])

	// 省略上标代表到 len(s) 结束
	fmt.Println("s[4:] ==", s[4:])
}

//3-13.range
func range_() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//3-16.map 的文法  map 的文法跟结构体文法相似，不过必须有键名
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func map_() {
	fmt.Println(m)
}

//3-16.函数值  函数也是值。他们可以像其他值一样传递，比如，函数值可以作为函数的参数或者返回值。
// 这个函数的参数是一个函数，相当于java中的回调
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func func_test() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}

//4-4.接口
