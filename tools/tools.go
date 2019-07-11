// tools project main.go
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	//what_make()
	runtimeUse()
}

// init 比main还早执行
func init() {
	//what_bytes2Str()
}

func what_make() {
	fmt.Println(make([]byte, 2))
	str := make([]byte, 3) //什么类型：=make(什么类型，长度)
	str[1] = 33
	fmt.Println(str)
}

//https://www.cnblogs.com/piperck/p/6471352.html
func what_timeFormat() {
	x := time.Date(2017, 02, 27, 17, 30, 20, 20, time.Local)
	fmt.Println(x.Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format("010215"))
}

//https://blog.csdn.net/a64180190/article/details/76528063
func what_bytes2Str() {
	var str1 string = "test"
	var data1 []byte = []byte(str1)
	fmt.Println(data1)

	var data2 [10]byte
	data2[0] = 'T'
	data2[2] = 'E'
	data2[4] = '1'
	var str2 string = string(data2[:])
	fmt.Println(str2)
}

func string_startWithEndWith() {
	fmt.Println(strings.HasPrefix("prefix", "pre"))
	fmt.Println(strings.HasPrefix("prefix", "pr"))
	fmt.Println(strings.HasSuffix("prefix", "fix"))
	fmt.Println(strings.HasSuffix("prefix", "i"))
	fmt.Println(strings.HasSuffix("prefix", "x"))
	fmt.Println(strings.HasSuffix("prefix", "X"))
}

func string_subString() {
	//包含中文的字符串
	s2 := "我是中国人"
	fmt.Println(s2[2:4])
	fmt.Println(string([]rune(s2)[2:4]))
}

//https://www.cnblogs.com/mafeng/p/7068837.html
func httpGet() {
	resp, err := http.Get("http://120.79.68.10:8080/gcluster/dashboard")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("connetcError:" + resp.Status)
	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("http://120.79.68.10:8080/gcluster/dashboard",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("connetcError:" + resp.Status)
	}

	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("http://120.79.68.10:8080/gcluster/dashboard",
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

//复杂的请求
func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://120.79.68.10:8080/gcluster/dashboard", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpDoParam() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://localhost:8081/t",
		strings.NewReader("a=cjb&b=什么既怕"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

var urlPath = "http://192.168.35.22:8080/api/v3/"
var authKey = "584167385b2e9"
var secretKey = "Mjg2NzI4Mzk5NTU3NjYxNTczNDMwNjM3NzAwODcwOTYzMjA"

//https://www.cnblogs.com/rinack/p/7568020.html
func httpBaseAuth() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", urlPath+"brokers", strings.NewReader(""))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", getBaseAuthHead(authKey, secretKey))

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
func getBaseAuthHead(authKey string, secretKey string) string {
	auth := authKey + ":" + secretKey
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	authHeader := "Basic " + string(encodedAuth)
	return authHeader
}

func DecHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

func Fill(sour string, fillStr string, size int, isLeft bool) string {
	if len(sour) == 0 {
		sour = ""
	}

	fillLen := size - len(sour)
	fill := ""

	for i := 0; i < fillLen; i = i + 1 {
		fill = fill + fillStr
	}
	if isLeft {
		return fill + sour
	} else {
		return sour + fill
	}
}

// hex string to bytes
func HexStringToBytes(s string) []byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}
	return bs
}

func SendSocket(address string, data []byte) []byte {
	bytedatas := make([][]byte, 2)
	bytedatas[0] = HexStringToBytes(Fill(DecHex(int64(len(data))), "0", 4, true))
	bytedatas[1] = []byte(data)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}

	_, err = conn.Write(bytes.Join(bytedatas, []byte("")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
	redata := readConn(conn, false)
	fmt.Fprintf(os.Stderr, "Fatal finish: %s", redata)
	conn.Close()
	return redata
}

// 接收数据
// 接收数据
func readConn(conn net.Conn, needLength bool) []byte {
	// read from the connection
	var lengthdata = make([]byte, 2)
	log.Println("start to read from conn")
	n, err := conn.Read(lengthdata)
	if err != nil {
		log.Println("conn read error:", err)
	} else {
		log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(lengthdata))
	}
	var receiveLength = ((int(lengthdata[0]) << 10) | int(lengthdata[1]))
	//receiveLength=11
	var buf = make([]byte, receiveLength)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("conn read error:", err.Error())
			return []byte("")
		} else {
			log.Printf("read %d bytes, content is:[%s]\n", n, hex.EncodeToString(buf)) //bytes-->16进制

		}
		break
	}

	if needLength {
		lenData := make([][]byte, 2)
		lenData[0] = lengthdata
		lenData[1] = buf
		returnData := bytes.Join(lenData, []byte(""))
		return returnData
	} else {
		return buf
	}
}

func RunSocketServer() {
	// lsof -i tcp:port    lsof -i tcp:12308
	addr := ":8086" //表示监听本地所有ip的8080端口，也可以这样写：addr := ":8080"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Printf(err.Error())
	}
	defer listener.Close()

	log.Printf(addr)

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Printf(err.Error())
		}
		go handle_conn(conn) //开启多个协程。
	}
}

//  处理socket连接
func handle_conn(conn net.Conn) {

	// 1.读取报文
	recvstr := readConn(conn, false)
	log.Printf("receive=====================================================\n")
	log.Printf("receive len:[%d] data:[%s]\n", len(recvstr), string(recvstr))
	log.Printf("============================================================\n")
	//// 2.处理报文
	//ifsucc:=mqttAdaptor.Publish("/groopush/serverToClient/command/N6NL00000004", recvstr)
	ifsucc := true
	log.Printf("send=====================================================\n")
	log.Printf("send len:[%d] data:[%s]\n", len(recvstr), ifsucc)
	log.Printf("=========================================================\n")
	data := HexStringToBytes("00023030")
	if !ifsucc {
		data = HexStringToBytes("00023131")
	}
	// 3.发送报文并关闭连接
	writeConn(conn, data)
	// 4.
	closeConn(conn)

}

// 发送数据
func writeConn(conn net.Conn, downmsg []byte) {
	//conn.Write(byteutil.HexStringToBytes(writeStr)) //通过conn的wirte方法将这些数据返回给客户端。
	conn.Write(downmsg) //通过conn的wirte方法将这些数据返回给客户端。

}

// 关闭连接
func closeConn(conn net.Conn) {
	time.Sleep(time.Minute)
	conn.Close() //与客户端断开连接。
}

// 打印什么系统
func runtimeUse() {
	fmt.Println(runtime.GOOS)
}
