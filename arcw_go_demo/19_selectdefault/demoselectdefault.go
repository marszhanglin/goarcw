package main

import "fmt"

/**
非阻塞通道 default
*/
func main() {

	message := make(chan string, 1)
	//message:=make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("receive:", msg)
	default:
		fmt.Println("nodata receive")
	}
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent:", msg)
	default:
		fmt.Println("nodata sent")
	}

	select {
	case msg := <-message:
		fmt.Println("received msg:", msg)
	case sig := <-signals:
		fmt.Println("received sig:", sig)
	default:
		fmt.Println("nodata")
	}
}
