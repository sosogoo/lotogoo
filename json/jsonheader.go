package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// Fruitbasket ...
type Fruitbasket struct {
	Name  string `json:"name"`
	Fruit []string `json:"fruit"`
	ID  int64 `json:"id"`
	Created time.Time `json:"created"`
}
//字符串数组
func jsonarry()  {
	jsondata:=[]byte(`{
		"name": "Standard",
        "fruit": [
             "Apple",
            "Banana",
            "Orange"
        ],
        "id": 999,
        "created": "2018-04-09T23:00:00Z"
	}`)
	var basket Fruitbasket
	if err:=json.Unmarshal(jsondata,&basket);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(basket)
}
// Fruit ...
type Fruit struct{
	Name string `json:"name"`
	PriceTag string `json:"priceTag"`
}
//Fruitdick ...
type Fruitdick struct {
	Name  string `json:"name"`
	Fruit Fruit `json:"fruit"`
	ID  int64 `json:"id"`
	Created time.Time `json:"created"`
}
//字典  内嵌对象
func jsondick() {
	jsondata := []byte(`
    {
        "name": "Standard",
        "fruit" : {"name": "Apple", "priceTag": "$1"},
        "def": 999,
        "created": "2018-04-09T23:00:00Z"
    }`)
	var basket Fruitdick
	if err:=json.Unmarshal(jsondata,&basket);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(basket)
}

// Fruitstring ...
type Fruitstring struct{
	Name string `json:"name"`
	PriceTag string `json:"priceTag"`
}
//Fruitdickstring ...
type Fruitdickstring struct {
	Name  string `json:"name"`
	Fruit []Fruitstring `json:"fruit"`
	ID  int64 `json:"id"`
	Created time.Time `json:"created"`
}
//字典数组 
func jsondickstring()  {
	jsondata := []byte(`
    {
        "name": "Standard",
        "fruit" : [{"name": "Apple", "priceTag": "$1"},{"name": "Pear","priceTag": "$1.5"}],
        "def": 999,
        "created": "2018-04-09T23:00:00Z"
    }`)
	var basket Fruitdickstring
	if err:=json.Unmarshal(jsondata,&basket);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(basket)
}

// Fruitkeydick ...
type Fruitkeydick struct{
	Name string `json:"name"`
	PriceTag string `json:"priceTag"`
}
//Fruitdickkey ...
type Fruitdickkey struct {
	Name  string `json:"name"`
	Fruit map[string]Fruitkeydick `json:"fruit"`
	ID  int64 `json:"id"`
	Created time.Time `json:"created"`
}
//动态key对象数据
func main()  {
	jsondata := []byte(`
    {
        "name": "Standard",
        "fruit" : {
			"1":{"name": "Apple", "priceTag": "$1"},
			"2":{"name": "Pear","priceTag": "$1.5"}
		    },
        "id": 999,
        "created": "2018-04-09T23:00:00Z"
    }`)
	var basket Fruitdickkey
	if err:=json.Unmarshal(jsondata,&basket);err!=nil{
		fmt.Println(err)
	}
	fmt.Println(basket)
}

// User  ....
type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}
func jsonuser() {

	p := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}
	//json编码
	byt, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	//json解析
	var ujson map[string]interface{}
	err = json.Unmarshal(byt, &ujson)
	if err != nil {
		panic(err)
	}
	fmt.Println("编码:", string(byt))
	fmt.Println("解析:", ujson)

}

//用空接口类型解析json数据
func basketjson() {
	jsondata := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)
	var v interface{}
	json.Unmarshal(jsondata, &v)
	data := v.(map[string]interface{})
	for k, v := range data {
		//类型断言
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array)")
			for i, u := range v {
				fmt.Println(" ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}

//用Decoder解析数据流(打开文件或者http请求体,他们都是io.Reader的实现)
func jsondecoder() {
	const jsonStream = `
    {"Name": "Ed", "Text": "Knock knock."}
    {"Name": "Sam", "Text": "Who's there?"}
    {"Name": "Ed", "Text": "Go fmt."}
    {"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}`
	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s:%s\n", m.Name, m.Text)
	}
}
