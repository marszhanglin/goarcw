package main

import (
	"fmt"
	"net/url"
)

/**
url
url 解析
*/
func main() {

	s := "postgres://User:pass@hot.com:8088/path?k=v#f"
	u, _ := url.Parse(s)
	fmt.Println(u.Scheme)
	fmt.Println(u.User.Username())
	fmt.Println(u.User.Password())
	fmt.Println(u.Host)
	fmt.Println(u.Hostname())
	fmt.Println(u.Port())
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)

}
