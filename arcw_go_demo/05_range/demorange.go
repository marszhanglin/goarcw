package main

import "fmt"

/**
range
遍历 数组 切片 字典
*/
func main() {
	// 切片 打印索引、值  计算所有元素和
	nums := []int{98, 67, 83}
	sum := 0
	for index, num := range nums {
		fmt.Println("index:", index, "value:", num)
		sum += num
	}
	fmt.Println("sum:", sum)

	// 遍历字典
	kvs := map[string]string{"张林": "xxg", "娟娟": "sgn"}
	for key, value := range kvs {
		fmt.Println("key:", key, "value:", value)
	}

}
