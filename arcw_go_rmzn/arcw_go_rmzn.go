package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
)

// 书籍：go入门指南
// pdf文件位置：MAC：/arcw/book/go入门指南
func main() {
	//regexp_repleace()
	getLocalIp()
}

// 4.9指针
func getzzdz() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	// *int 地址，
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	s := "good bye"
	var p *string = &s                            // &s表示s的地址
	*p = "ciao"                                   // p:地址   *p:string
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

}

// 6.2.1 按值传递参数（call by value）  按引用传递参数（call by reference）
func cbvcbr() {
	a := "1"
	cvalue(a)
	fmt.Println(a)
	czzvalue(&a)
	fmt.Println(a)
}

func cvalue(a string) {
	a = "2"
	fmt.Println(a)
}

func czzvalue(a *string) {
	*a = "3"
	fmt.Println(a)
}

// 6.4 defer 和追踪  相当于java的finally
func defert1() {
	i := 0
	fmt.Println("1")
	defer fmt.Println("defer like finally", i)
	i++
	fmt.Println("2")

}

func defert2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("多个defer时以栈的方式执行", i)
	}

}

// 6.4 追踪  相当于java的finally
func trace(s string) string {
	fmt.Println("in------------------------------------>", s)
	return s
}
func out(s string) {
	fmt.Println("out------------------------------------>", s)
}

func defert3() {
	defer out(trace("defert3"))
	fmt.Println("do something...")
}

// 6.11 计算函数执行时间
func useTimet1() {
	start := time.Now()
	fmt.Println("我的天", "这么快", "这么快", "这么快", "这么快")
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf(" took this amount of time: %s\n", delta)
}

// 切片 0 <= len(s) <= cap(s)
// 7.5 切片的复制与追加
func copytappendt() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10, 12)
	n := copy(sl_to, sl_from)
	fmt.Print(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

}

// 8.1 声明、初始化和 make
func makeInit() {
	var mapInit1 map[string]int
	mapInit1 = make(map[string]int)
	mapInit1["1"] = 2

	mapInit2 := map[string]int{"2": 3}
	fmt.Printf("Map  is: %d\n", mapInit1["1"])
	fmt.Printf("Map  is: %d\n", mapInit2["2"])

	fmt.Printf("Map  is: %d\n", mapInit1["3"])

	value3, ok := mapInit2["3"]
	fmt.Println(value3, ok)
	value2, ok := mapInit2["2"]
	fmt.Println(value2, ok)

	// key values
	mapInit3 := make(map[string][]int)
	mapInit3["2"] = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Map  is: %d\n", mapInit3["2"])
}

// 8.3 map for-range 的配套用法
// 注意 map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。
func mapRange() {
	mapdaata := make(map[string]int)
	mapdaata["b"] = 2
	mapdaata["a"] = 1
	mapdaata["c"] = 3
	mapdaata["d"] = 4
	mapdaata["e"] = 5
	for key, value := range mapdaata {
		fmt.Println(key, value)
	}
}

//8.4 map 类型的切片
func mapsRange() {
	items1 := make([]map[int]int, 5)
	for i := range items1 {
		items1[i] = make(map[int]int, 1)
		items1[i][1] = 3
	}
	fmt.Printf("Version A: Value of items: %v\n", items1)

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 3
	}
	fmt.Printf("Version A: Value of items: %v\n", items2)

}

//
//strings : 提供对字符串的操作。
//strconv : 提供将字符串转换为基础类型的功能。
//unicode : 为 unicode 型的字符串提供特殊的功能。
//regexp : 正则表达式功能。
//bytes : 提供对字符型分片的操作。
//index/suffixarray : 子字符串快速查询。
//math - math/cmath - math/big - math/rand - sort :
//	math : 基本的数学函数。
//math/cmath : 对复数的操作。
//math/rand : 伪随机数生成。
//sort : 为数组排序和自定义集合。
//math/big : 大数的实现和计算。
//container - /list-ring-heap : 实现对集合的操作。
//list : 双链表。
//ring : 环形链
//time - log :
//time : 日期和时间的基本操作。
//log : 记录程序运行时产生的日志,我们将在后面的章节使用它。
//encoding/json - encoding/xml - text/template :
//encoding/json : 读取并解码和写入并编码 JSON 数据。
//encoding/xml :简单的 XML1.0 解析器,有关 JSON 和 XML 的实例请查阅第 12.9/10 章节。
//text/template :生成像 HTML 一样的数据与文本混合的数据驱动模板（参见第 15.7 节）。
//net - net/http - html :（参见第 15 章）
//net : 网络数据的基本操作。
//http : 提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。
//html : HTML5 解析器。
//runtime : Go 程序运行时的交互操作，例如垃圾回收和协程创建。
//reflect : 实现通过程序运行时反射，让程序操作任意类型的变量。
//archive/tar 和 /zip-compress ：压缩(解压缩)文件功能。
//fmt - io - bufio - path/filepath - flag :fmt : 提供了格式化输入输出功能。
//io : 提供了基本输入输出功能，大多数是围绕系统功能的封装
//bufio : 缓冲输入输出功能的封装。
//path/filepath : 用来操作在当前系统中的目标文件名路径。
//flag : 对命令行参数的操作。strings - strconv - unicode - regexp - bytes

// 9.2 regexp 包 正则隐藏金额
func regexp_repleace() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
}

//接口==========================================来个会吃饭的人
type DoSomething interface {
	eat() string
	sleep() string
}

// 人类
type Human struct {
	name string
}

// 这个人实现是吃
func (hm *Human) eat() string {
	return hm.name + "吃了五斤肉"
}

// 这个人实现是睡
func (hm *Human) sleep() string {
	return hm.name + "吃了五斤肉"
}

// 来个会吃饭的人
func hotel(xiaowan DoSomething) {
	fmt.Print(xiaowan.eat())
	//fmt.Print(xiaowan.sleep())
}

//接口==========================================来个会吃饭的人

// ==============================================获取本机ip
func getLocalIp() (ip string) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if strings.Contains(ipnet.IP.String(), "192.168") {
					return ipnet.IP.String()
				}
			}

		}
	}
	return ""
}
