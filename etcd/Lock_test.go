package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	// 客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"},
		DialTimeout: 5 * time.Second,
	}
	//str := []string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}
	//client, _ := newClient(str)
	// 建立连接
	client1, err := clientv3.New(config)
	//client1, err := newClient(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 1. 上锁并创建租约
	lease := clientv3.NewLease(client1)
	leaseGrantResp, err := lease.Grant(context.TODO(), 5)
	if err != nil {
		panic(err)
	}
	leaseId := leaseGrantResp.ID
	// 2 自动续约
	// 创建一个可取消的租约，主要是为了退出的时候能够释放
	ctx, cancelFunc := context.WithCancel(context.TODO())
	// 3. 释放租约
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)
	keepRespChan, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		panic(err)
	}
	// 续约应答

	go func() {
		for {
			select {
			case keepResp := <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else { // 每秒会续租一次, 所以就会受到一次应答
					fmt.Println("收到自动续租应答:", keepResp.ID)
				}
			}
		}
	END:
	}()
	// 1.3 在租约时间内去抢锁（etcd 里面的锁就是一个 key）
	kv := clientv3.NewKV(client1)
	// 创建事务
	txn := kv.Txn(context.TODO())
	//if 不存在 key，then 设置它，else 抢锁失败
	txn.If(clientv3.Compare(clientv3.CreateRevision("lock"), "=", 0)).
		Then(clientv3.OpPut("lock", "g", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet("lock"))
	// 提交事务
	txnResp, err := txn.Commit()
	if err != nil {
		panic(err)
	}
	if !txnResp.Succeeded {
		fmt.Println("锁被占用:", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}
	// 抢到锁后执行业务逻辑，没有抢到退出
	fmt.Println("处理任务")
	time.Sleep(5 * time.Second)
}
