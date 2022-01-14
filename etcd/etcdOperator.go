package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type Cli struct {
	cli3 clientv3.Client
	// cli3 *clientv3.Client,看一看区别
}

func newClient1(str []string) (Cli, error) {
	cli, err := clientv3.New(clientv3.Config{
		//Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		Endpoints:   str,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close() // 该把这一步下放到其他函数里，入put，get等
	return Cli{cli3: *cli}, err
}
func newClient2(str1, str2, str3 string) (Cli, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{str1, str2, str3},
		//Endpoints: []string{str1, str2, str3},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println("ERRR HAPPEN")

	}
	//defer cli.Close() // 该把这一步下放到其他函数里，入put，get等
	return Cli{cli3: *cli}, err
}
func newClient(str []string) (Cli, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: str,
		//Endpoints: []string{str1, str2, str3},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println("ERRR HAPPEN")

	}
	//defer cli.Close() // 该把这一步下放到其他函数里，入put，get等
	return Cli{cli3: *cli}, err
}
func (c *Cli) put(key string, value string) error {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err := c.cli3.Put(ctx, key, value)
	//cancel() // cancel的作用？
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return err
	}
	//defer c.cli3.Close()
	return err
}
func (c *Cli) get(key string) []byte { // read data from etcd
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	resp, err := c.cli3.Get(ctx, key)
	//cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		panic(err)
		return nil
	}
	var value []byte
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		value = ev.Value
	}
	return value
}
func (c *Cli) delete(key string) error {
	// 获取上下文，设置请求超时时间为5秒
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 删除key="/tizi365/url" 的值
	_, err := c.cli3.Delete(ctx, key)

	if err != nil {
		panic(err)
	}
	return err
}
func (c *Cli) getWithPrefix(prefix string) []byte {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	resp, err := c.cli3.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	var value []byte
	// 遍历查询结果
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		value = ev.Value
	}
	return value
}
func (c *Cli) deleteWithPrefix(prefix string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err := c.cli3.Delete(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}
	return err
}
func (c *Cli) watchWithPrefix(prefix string) {
	// 监控以key为前缀的所有key的值
	rch := c.cli3.Watch(context.Background(), prefix, clientv3.WithPrefix())
	// 通过channel遍历key的值的变化
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
func (c *Cli) lease(key, value string, ttl int64) {
	// 创建一个ttl秒的租约
	resp, err := c.cli3.Grant(context.TODO(), ttl)
	if err != nil {
		panic(err)
	}

	// ttl秒钟之后, key 这个key就会被移除
	_, err = c.cli3.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}
}
func (c *Cli) keepAlive(key, value string, ttl int64) {
	resp, err := c.cli3.Grant(context.TODO(), ttl)
	if err != nil {
		panic(err)
	}

	_, err = c.cli3.Put(context.TODO(), key, value, clientv3.WithLease(resp.ID))
	if err != nil {
		panic(err)
	}

	// the key 'foo' will be kept forever
	ch, kaerr := c.cli3.KeepAlive(context.TODO(), resp.ID)
	if kaerr != nil {
		panic(err)
	}
	for {
		ka := <-ch
		fmt.Println("ttl:", ka.TTL)
	}
}
func concurrency() {

}
func NewLease(c *Cli) clientv3.Lease { // 待议

	p := &c.cli3
	return clientv3.NewLeaseFromLeaseClient(clientv3.RetryLeaseClient(p), p, 5*time.Second)
}
