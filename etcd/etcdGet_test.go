package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd client put/get demo
// use etcd/clientv3

func Test(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "test", "value")
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "test")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
}

func TestPut(t *testing.T) {
	var str = []string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}
	//client, _ := newClienttest("10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379")
	client, _ := newClient(str)
	defer client.cli3.Close()
	client.put("Master", "10.176.34.179")

	resp := client.get("test")
	fmt.Printf("%s\n", resp)
}

func TestGet(t *testing.T) {
	var str = []string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}
	//client, _ := newClienttest("10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379")
	client, _ := newClient(str)
	defer client.cli3.Close()
	//client.put("test","value3")

	resp := client.get("Master")
	fmt.Printf("%s\n", resp)
	for _, ev := range resp {
		fmt.Printf(string(ev))
	}
}
