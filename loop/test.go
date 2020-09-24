package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/jmoiron/sqlx"
	// "io/ioutil"
	// "time"
	// "github.com/gin-gonic/gin"
)

// type Pool struct {
// 	work chan func()   //任务
// 	sem  chan struct{} //数量
// }

// //定义一个new函数  ,创建一个协程池对象
// func New(size int) *Pool {
// 	return &Pool{
// 		// work: make(chan func()), //无缓冲通道
// 		// sem:make(chan struct{},size) //有缓冲通道
// 	}
// }

// //
// func (p *Pool) NewTask(task func()) {
// 	select {
// 	case p.work <- task:
// 	case p.sem <- struct{}{}:
// 		go p.worker(task) //开启一个协程
// 	}
// }
// func (p *Pool) worker(task func()) {
// 	defer func() { <-p.sem }()
// 	for {
// 		task()
// 		task = <-p.work
// 	}
// }
// func loop() {
// 	startTime := time.Now()
// 	for i := 1; i < 100000; i++ {
// 		fmt.Printf("%d\n", i)
// 	}
// 	endTime := time.Since(startTime)
// 	fmt.Println("用时:", endTime)
// }
//全局变量m
// var m = 100 //初始化变量
func foo() (int, string) {

	fmt.Println("123")
	return 10, "你好"
}

/*1.函数外的每个语句都要以关键字开始
  2.:=不能使用在函数外
  3._多用于占位,表示忽略值
*/

func lin() {
	// x, _ := foo()
	// _, y := foo()
	// g, err := foo()
	// fmt.Println("x=", x)
	// fmt.Println("y=", y)
	// fmt.Println("z=", g, err)
	// n := 10
	// m := 200 //此处声明局部变量
	// fmt.Println(m, n)
	// go loop() //开启一个线程
	// go loop() //开启一个线程
	// fmt.Printf("已经开始了")
	// time.Sleep(time.Second)
	// pool:=New(2)
	// for i:=1;i<5;i++{
	// 	pool.NewTask(func(){
	// 		time.Sleep(2 * time.Second)
	// 		fmt.Println(time.Now())
	//     })
	//    }
	//保证所有的协程都执行完毕
	// time.Sleep(5 * time.Second)
	// var city [5]string
	// city[0] = "北京"
	// city[1] = "上海"
	// city[2] = "广州"
	// city[3] = "深圳"
	// city[4] = "濮阳"
	// othercity := [...]string{"北京", "上海", "广州", "深圳", "贵阳"}
	//创建多维数组
	// var multi [2][3]int
	// for j := 0; j < 2; j++ {
	// 	for i := 0; i < 3; i++ {
	// 		multi[j][i] = j + i
	// 	}
	// }
	// fmt.Println(city[0], city[1], city[2], city[3], city[4])
	// fmt.Println("以不同的方式打印数组")
	//数字代表的是索引值
	// fmt.Println(othercity[1:])
	// fmt.Printf("%s\n", othercity)
	// fmt.Printf("%q\n", othercity)
	//创建切片
	// fmt.Println("用make创建切片")
	// cities := make([]string, 3)
	// other_cities := []string{}
	// cities[0] = "上海"
	// cities[1] = "深圳"
	// cities[2] = "温州"
	// fmt.Println(cities)
	// other_cities = append(other_cities, cities...)
	// fmt.Println(other_cities)
	// //拷贝切片
	// copy_cities := make([]string, len(other_cities))
	// copy(copy_cities, other_cities)
	// fmt.Println("copy:", copy_cities)

	//map是一种无序键值对的集合
	//指针类型  指针是值内存地址存放的地方 由*定义 根据数据类型指定指针
	//*号可以指定一个变量作为指针   &运算符可以获取变量的地址
	// var age int = 99 /*声明实际变量*/
	// var ip *int      /*声明指针变量*/
	// ip = &age        /*指针变量的存储地址*/
	// //指针变量的存储地址
	// fmt.Println("指针变量的存储地址:", ip)
	// //访问指针变量的值
	// fmt.Println("指针变量的值:", *ip, age)

	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() //监听 并在0.0.0.0:8080上启动服务
	//打开文件,文件路径相对于GOPATH开始,或者写全路径(/Users/xxx/Go/src/file.txt)
	// file, err := ioutil.ReadFile("D:/aabbcc.txt")
	// if err != nil {
	// 	fmt.Println("文件打开失败,原因是:", err)
	// }

	// fmt.Printf("%s", string(file))
	// str := "hello沙河县小吃街"
	// // s := []rune(str)
	// count := 0
	// for _, r := range str {
	// 	// fmt.Println(len(string(r)))
	// 	// fmt.Printf("%v(%c\n)", r, r)
	// 	if len(string(r)) >= 3 {
	// 		count++
	// 	}
	// }
	// fmt.Println("汉字的数量:", count)
	// gotodemo2()
	// breakdemo1()
	// continuedemo()
	// ergodic()
	// twoarray()
	// mapdemo()

}

func gotodemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakFlag = true
				break
			}
			fmt.Printf("%v,%v\n", i, j)
		}
		if breakFlag {
			break
		}
	}

}

//使用goto语句
func gotodemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				//设置退出标签
				goto breakTag
			}
			fmt.Printf("%v,%v\n", i, j)
		}
	}
	// return
	//标签
breakTag:
	fmt.Println("结束for循环")
}
func breakdemo1() {
BRE:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break BRE
			}
			fmt.Printf("%v,%v\n", i, j)
		}
	}
}
func continuedemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v,%v\n", i, j)
		}
	}
}

//遍历数组有两种方法
func ergodic() {
	var a = [...]int{1, 3, 5, 7, 8}
	//方式一 for遍历
	sum := 0
	sum1 := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]

	}
	fmt.Println(sum)
	//方式二 for range遍历
	for _, value := range a {
		sum1 += value

	}
	fmt.Println(sum1)
}
func twoarray() {
	var a = [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(a); i++ {
		// fmt.Fprintln(a[i])
		Num1 := a[i]
		for j := i + 1; j < len(a); j++ {
			// fmt.Println(a[j])
			Num2 := a[j]
			Sum := Num1 + Num2
			if Sum == 8 {
				fmt.Println(i, j)
			}

		}

	}
}

func mapdemo() {
	sremap := make(map[string]int)
	sremap["张三"] = 30
	sremap["王四"] = 90
	sremap["小明"] = 100
	for k := range sremap {
		fmt.Println(k)
	}
}
func test() {
	for i := 1; i <= 9; i++ {
		for a := 9; a > i; a-- {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf(" *")
		}
		fmt.Println("")
	}
	for b := 8; b >= 1; b-- {
		for c := 9; c > b; c-- {
			fmt.Printf(" ")
		}
		for d := 1; d <= b; d++ {
			fmt.Printf(" *")
		}
		fmt.Println("")
	}

}

var wg sync.WaitGroup

func hello1() {
	for i := 0; i < 10000; i++ {
		// defer wg.Done() //goroutine结束就登记-1
		fmt.Println("id1", i)
	}
}
func hello2() {
	for i := 0; i < 10000; i++ {
		// defer wg.Done() //goroutine结束就登记-1
		fmt.Println("id2", i)
	}
}

func gor() {
	starttime := time.Now()

	// wg.Add(1) //启动一个goroutine就登记+1
	runtime.GOMAXPROCS(2) //设置核心
	go hello1()
	go hello2()

	time.Sleep(time.Second)
	// wg.Wait() //等待所有登记的goroutine都结束
	endtime := time.Since(starttime)
	fmt.Println("用时:", endtime)
	// var a []string
	// var b = []int{}
	// var c = []bool{false, true}
	// var d = []bool{true, false}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(a == nil)
	// fmt.Println(b == nil)
	// fmt.Println(c == nil)
	// fmt.Println(d == nil)

}

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func Qrcode() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET"><input maxLength=1024 size=70
name=s value="" title="Text to QR Encode"><input type=submit
value="Show QR" name=qr>
</form>
</body>
</html>
`

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func pass() {
	// var ch1 chan int
	// ch1 := make(chan int)    //这种是创建无缓冲通道  必须有接收者才能发送
	ch1 := make(chan int, 2) //创建通道容量为2的有缓冲通道
	// ch2 := make(chan bool)
	// go recv(ch1) //启用一个通道接收
	for i := 0; i < 3; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:
		}
		// ch1 <- i //将i发送到通道ch1中
		// x := <-ch1 //从ch1中接收值并赋值给变量x

	}
	close(ch1)
	fmt.Println("发送成功")
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect fail,err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

//查询单行数据
func queryRowDemo() {
	sqlStr := "SELECT * from name; "
	var u int
	err := db.Get(&u, sqlStr)
	if err != nil {
		fmt.Println("get fail,err:%v\n", err)
		return
	}
	fmt.Println(u)
}

func star() {

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print("*")
			fmt.Print("\n")

		}
		fmt.Println("")
	}

}

//做一个随机数
func RandDemo() {
	//若想做一个真正的随机数 要种子  seed()种子是默认是1
	//rand.Seed(1)  设置随机数种子  保证每次随机数据是不重复
	rand.Seed(time.Now().UnixNano())

	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = rand.Intn(1000)
	}
	// fmt.Println(time.Now().UnixNano())
	fmt.Printf("%.3d\n", b)
}

func SliceDemo() {
	var s1 []int
	s2 := []int{}
	s3 := make([]int, 0)
	var s4 []int = make([]int, 0, 0)
	s5 := []int{1, 2, 3, 4}
	fmt.Println(s1, s2, s3, s4, s5)
	fmt.Println("----------------------")
	//从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int //初始化[]
	//左包右不包
	s6 = arr[:len(arr)-1] //开始赋值
	fmt.Println(s6)
	fmt.Println("----------------------")
	//通过make创建切片
	// 格式 var slice []type=make([]type,len)
	// 格式 slice :=make([]type,len)
	// 格式 slice :=make([]type,len,cap)
	var slice0 []int = make([]int, 10)
	slice1 := make([]int, 10)
	slice3 := make([]int, 10, 10)
	fmt.Println(slice0, slice1, slice3)
	fmt.Println("----------------------")
	//
	s := []int{1, 2, 3, 4, 5}
	p := &s[2] //获取底层数组元素指针
	*p += 100
	fmt.Println(s, p, &p, *p)
	fmt.Println("----------------------")
	//超出原slice.cap限制,就会重新分配底层数组 即使原数组并未填满
	data := [...]int{1, 2, 3, 4, 5, 10: 0}
	s9 := data[:2:3]
	fmt.Println(s9)
	fmt.Println("----------------------")
	h := make([]int, 0, 1)
	c := cap(h)
	for i := 0; i < 9; i++ { //超过cap限制  重新分配内存
		h = append(h, i)
		c = cap(h)
	}
	fmt.Println(h)
	fmt.Println(c)
}
func demolice() {
	//拷贝切片
	// a := []int{1, 2, 3}
	// b := make([]int, 10)
	// fmt.Printf("a:%v\nb:%v\n", a, b)
	// copy(b, a)
	// fmt.Printf("a:%v\nb:%v\n", a, b)
	// a = append(a, b...)
	// fmt.Printf("a:%v\nb:%v\n", a, b)
	// a = append(b, 4, 5, 6)
	// fmt.Printf("a:%v\nb:%v\n", a, b)
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[:8]
	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))
	//将切片或者数组转成字符串
	e := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1)
	fmt.Println(e)
	fmt.Println("----------------------")
	m := make([]byte, 200)
	ptr := unsafe.Pointer(&m[0])
	fmt.Println(ptr)
}

func ReadSource(reader io.Reader) <-chan int {
	out := make(chan int)

	go func() {
		buffer := make([]byte, 8)

		for {
			n, err := reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
		// fmt.Println(buffer)
	}

}

func RandomSend(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
