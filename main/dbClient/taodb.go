package main

import (
	"fmt"
	"github.com/tsandl/TYDB/client"
	"github.com/tsandl/TYDB/file/operator"
	"github.com/tsandl/TYDB/log"
	"io/ioutil"
	"time"
)

// GLOBALS
var (
	opts *Options
)

func main() {
	opts = parseArgs()

	// init logging
	log.LogTo(opts.logto, opts.loglevel)
	client := client.New(opts.ip + opts.port)

	//for i := 0; i < 100; i++ {
	//	client.Set(fmt.Sprintf("hello%d", i), []byte(fmt.Sprintf("Hello World!%d", i)))
	//}
	//l, data3 := operator.ReadFile("F:\\huajianmodel\\data_model\\picture\\11.jpg")
	//fmt.Println(l,data3)
	//fmt.Println("function test..........")
	//defer func() {
	//	_, err := clients.CloseDB()
	//	if err == nil {
	//		fmt.Printf("Success close db")
	//	} else {
	//		fmt.Printf("fail close db")
	//	}
	//
	//}()
	//func() {
	//	_,err:= clients.OpenDB("F:\\data\\storage\\db6")
	//	if err == nil{
	//		fmt.Println("make success")
	//	}
	//	data, err := ioutil.ReadFile("F:\\huajianmodel\\data_model\\picture\\11.jpg")
	//	clients.Set("test", data)
	//	clients.Set("test", []byte("test"))
	//	fmt.Println(clients.Get("test"))
	//
	//
	//}()
	client.OpenDB("F:\\data\\storage\\db8")
	data4, _ := client.Get("picture2")
	operator.WriteFile("F:\\code\\go_code\\db1\\file\\result\\picture14.jpg", data4)
	client.CloseDB()
	data3, _ := client.Get("picture3")
	operator.WriteFile("F:\\code\\go_code\\db1\\file\\result\\picture13.jpg", data3)

	data, err := ioutil.ReadFile("F:\\huajianmodel\\data_model\\picture\\1.jpg")
	if err != nil {
		fmt.Println("errors happen")
	} else {
		fmt.Println(len(data))
	}

	client.Set("bigPicture", data)
	start := time.Now()
	for i := 0; i < 10; i++ {
		client.Set(fmt.Sprintf("picture%d", i), data)
	}
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
	data2, _ := client.Get("picture2")
	operator.WriteFile("F:\\code\\go_code\\db1\\file\\result\\picture999.jpg", data2)
<<<<<<< HEAD

=======
	// 需要將上傳的所有文件名記錄一下，這樣從minio中讀取時知道讀取什麽
>>>>>>> version_1.0.0
	//for i := 0; i < 100; i++ {
	//	value, _ := client.Get(fmt.Sprintf("hello%d", i))
	//	log.Info("get key:hello%d,value=%s\n", i, string(value[:]))
	//	//fmt.Println("hello")
	//	//fmt.Println(value[:])
	//}
	//
	//ctx, _ := client.Prefix("hello")
	//if len(ctx) == 0 {
	//	log.Info("ctx is null")
	//} else {
	//	fmt.Println("cit is not null")
	//	data := make(map[string]string)
	//	err := json.Unmarshal(ctx, &data)
	//	if err != nil {
	//		log.Error("json error:", err)
	//	}
	//	if len(data) > 0 {
	//		for key, value := range data {
	//			log.Info("pre.key=%s,%s\n", key, value)
	//		}
	//	}
	//}
	//ctx, _ = client.PrefixOnlyKey("hello")
	//if len(ctx) == 0 {
	//	log.Info("ctx is null")
	//} else {
	//	data := make([]string, 0)
	//	err := json.Unmarshal(ctx, &data)
	//	if err != nil {
	//		log.Error("json error:", err)
	//		return
	//	}
	//	if len(data) > 0 {
	//		log.Info("data.len=%d", len(data))
	//		for i, key := range data {
	//			log.Info("pre.i=%d,key=%s", i, key)
	//			value, _ := client.Get(key)
	//			log.Info("getValue=%s", value[:])
	//		}
	//	}
	//
	//}

}
