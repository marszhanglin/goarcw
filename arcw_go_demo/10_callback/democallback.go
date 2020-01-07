package _0_callback

import "fmt"

/**
callBack
*/
func main() {
	x, y := 1, 2
	fmt.Println(test(x, y, func(a, b int) int {
		fmt.Println("do callBack")
		return a + b
	}))

}

type CallbackInterFacefunc func(x, y int) int

func test(x, y int, callBackfunc CallbackInterFacefunc) int {
	fmt.Println("do test")
	return callBackfunc(x, y)
}
