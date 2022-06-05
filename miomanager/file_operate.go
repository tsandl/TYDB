package minioServer

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read file by filPath and return []byte

func readFile(filePath string) []byte {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return content
}

//
func readFile1(filePath string) (*os.File, int64) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//n, err := minioClient.PutObject("db5", "000005.ldb", file, fileStat.Size(), minio.PutObjectOptions{ContentType:"application/octet-stream"})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Successfully uploaded bytes: ", n)

	return file, fileStat.Size()
}
