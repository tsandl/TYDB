package minio

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestNewMinIO(t *testing.T) {
	minClient := NewMinIO("127.0.0.1:9000", "admin", "admin123456")
	fmt.Println(minClient.mClient.BucketExists("db1"))
}
func TestMakeBucket(t *testing.T) {
	minClient := NewMinIO("127.0.0.1:9000", "admin", "admin123456")
	minClient.makeBucket("dbtest")
}

func TestRemoveBucket(t *testing.T) {
	minClient := NewMinIO("127.0.0.1:9000", "admin", "admin123456")
	minClient.removeBucket("dbtest")
}

func TestPutObject(t *testing.T) {
	minClient := NewMinIO("127.0.0.1:9000", "admin", "admin123456")
	bucketName := "dbtest"
	objectName := "test1.png"
	filePath := "C:\\Users\\tsxqc\\Pictures\\Camera Roll\\haha.png"
	file, filestat := readFile1(filePath)
	defer file.Close() // 此步来源于readFile1()，若在readfile1中使用file.close()，数据是没办法传到这边来，因此需要在此处关闭文件。
	n, err := minClient.putObject(bucketName, objectName, file, filestat)
	fmt.Println(n, err)
}
func TestGetObject(t *testing.T) {
	minClient := NewMinIO("127.0.0.1:9000", "admin", "admin123456")
	object, err := minClient.getObject("demo", "last_love.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	localFile, err := os.Create("./local-file.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err = io.Copy(localFile, object); err != nil {
		fmt.Println(err)
		return
	}
}
