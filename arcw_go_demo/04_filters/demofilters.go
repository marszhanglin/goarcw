package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
json
*/
func main() {
	scanFile("/Users/marszhang/log/log", "")
}

func scanFile(filePath string, cron string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineData := scanner.Text()
		if strings.Contains(lineData, cron) {
			fmt.Println(lineData)
		}
	}
}
