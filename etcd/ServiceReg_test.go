package etcd

import (
	"fmt"
	"testing"
)

func TestReg(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/master", "10.176.34.178")
	select {}
}
func TestReg1(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave1", "10.176.34.130")
	select {}
}
func TestReg2(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave2", "10.176.34.132")
	select {}
}
func TestReg3(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave3", "10.176.34.135")
	select {}
}
func TestReg4(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave4", "10.176.34.179")
	select {}
}
func TestReg5(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave5", "10.176.34.181")
	select {}
}
func TestReg6(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave6", "10.176.34.182")
	select {}
}
func TestReg7(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave7", "10.176.34.183")
	select {}
}
func TestReg8(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave8", "10.176.34.188")
	select {}
}
func TestReg9(t *testing.T) {
	ser, _ := NewServiceReg([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"}, 5)
	ser.PutService("/node/slave9", "10.176.34.191")
	select {}
}

var addr []string

func TestDis(t *testing.T) {
	cli, _ := NewClientDis([]string{"10.176.34.134:2379", "10.176.34.135:2379", "10.176.34.140:2379"})
	addr, _ = cli.GetService("/node")
	for _, s := range addr {
		fmt.Println(s)
	}
	select {}
}
