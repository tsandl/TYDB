package main

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"io/ioutil"
	"os"
	"testing"
)

func TestConect(t *testing.T) {
	ssl := false
	// init minio client object.
	minioClient, err := minio.New("127.0.0.1:9000", "admin", "admin123456", ssl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("prepare created mybucket.")
	err = minioClient.MakeBucket("db5", "cn-north-1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully created mybucket.")
	found, err := minioClient.BucketExists("db3")
	if err != nil {
		fmt.Println(err)
		return
	}
	if found {
		fmt.Println("Bucket found")
	}

}

func TestPut(t *testing.T) {
	ssl := false
	// init minio client object.
	minioClient, err := minio.New("127.0.0.1:9000", "admin", "admin123456", ssl)

	filePath := "F:\\work\\testdb1\\db2\\000005.ldb"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	n, err := minioClient.PutObject("db5", "000005.ldb", file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", n)
}

func TestReadFile(t *testing.T) {
	// 读取当前目录中的所有文件和子目录
	filePath := "F:/data/storage/db4"
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		println(file.Name())
		// need to be finished read file by file path.
	}
}
