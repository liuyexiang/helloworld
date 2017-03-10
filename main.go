package main

import (
	"encoding/json"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"os"
	"sync"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

}

func testchan() {
	var waitGroup sync.WaitGroup // 用于等待一组操作执行完毕的同步工具。
	waitGroup.Add(3)

	numberChan1 := make(chan int64, 100) // 数字通道1。
	numberChan2 := make(chan int64, 100) // 数字通道2。
	numberChan3 := make(chan int64, 100) // 数字通道3。

	for i := 0; i < 100; i++ {
		numberChan1 <- rand.Int63n(100)
	}
	close(numberChan1)
	go func() { // 数字过滤装置1。
		for n := range numberChan1 { // 不断的从数字通道1中接收数字，直到该通道关闭。
			if n%2 == 0 { // 仅当数字可以被2整除，才将其发送到数字通道2.
				numberChan2 <- n
			} else {
				fmt.Printf("Filter %d. [filter 2]\n", n)
			}
		}
		close(numberChan2) // 关闭数字通道2。
		waitGroup.Done()   // 表示一个操作完成。
	}()

	go func() { // 数字过滤装置1。
		for n := range numberChan2 { // 不断的从数字通道1中接收数字，直到该通道关闭。
			if n%5 == 0 { // 仅当数字可以被2整除，才将其发送到数字通道2.
				numberChan3 <- n
			} else {
				fmt.Printf("Filter %d. [filter 5]\n", n)
			}
		}
		close(numberChan3) // 关闭数字通道2。
		waitGroup.Done()   // 表示一个操作完成。
	}()

	go func() { // 数字输出装置。
		for n := range numberChan3 { // 不断的从数字通道3中接收数字，直到该通道关闭。
			fmt.Println(n) // 打印数字。
		}
		waitGroup.Done() // 表示一个操作完成。
	}()

	waitGroup.Wait() // 等待前面那组操作（共3个）的完成。
}

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func (s *Student) ShowStu() {
	fmt.Println("show Student :")
	fmt.Println("\tName\t:", s.Name)
	fmt.Println("\tAge\t:", s.Age)
	fmt.Println("\tGuake\t:", s.Guake)
	fmt.Println("\tPrice\t:", s.Price)
	fmt.Printf("\tClasses\t: ")
	for _, a := range s.Classes {
		fmt.Printf("%s ", a)
	}
	fmt.Println("")
}

func testjsonmarchalunmarshal() {
	st := &Student{
		"Xiao Ming",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}
	fmt.Println("before JSON encoding :")
	st.ShowStu()

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println(b)
		fmt.Println("string b:")
		fmt.Println(string(b))
		fmt.Println("**************************")
	}
	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))
	strData := <-ch
	fmt.Println("--------------------------------")
	stb := &Student{}
	stb.ShowStu()
	err = json.Unmarshal([]byte(strData), &stb)
	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		fmt.Println("Unmarshal success")
		stb.ShowStu()
	}
}

var (
	levelFlag = flag.Int("level", 0, "级别")
	bnFlag    int
)

func init() {
	flag.IntVar(&bnFlag, "bn", 3, "份数")
}

func testflag() {

	flag.Parse()
	count := len(os.Args)
	fmt.Println("参数总个数:", count)

	fmt.Println("参数详情:")
	for i := 0; i < count; i++ {
		fmt.Println(i, ":", os.Args[i])
	}

	fmt.Println("\n参数值:")
	fmt.Println("级别:", *levelFlag)
	fmt.Println("份数:", bnFlag)
}

func testSlice() {
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println("slice cap", cap(mySlice))

}

type personInfo struct {
	ID      string
	Name    string
	Address string
}

func tryMap() {
	var personDB map[string]personInfo
	personDB = make(map[string]personInfo)

	personDB["1"] = personInfo{"1", "lily", "5-607"}
	personDB["2"] = personInfo{"2", "lucy", "4-908"}

	value, ok := personDB["2"]

	if ok {
		fmt.Println("found id 2", value.ID, value.Address, value.Name)
	} else {
		fmt.Println("Did not find person with ID 2")
	}
}

func tryRSA() {
	var bits int
	flag.IntVar(&bits, "b", 2048, "密钥长度，默认为1024位")
	fmt.Println(bits)
}
