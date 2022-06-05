package minioServer

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"testing"
)

func TestFile(t *testing.T) {
	filePath := "C:\\Users\\tsxqc\\Pictures\\Camera Roll\\haha.png"
	file, fileStat := readFile1(filePath)
	// readFile1() 中 defer file.Close() 注释掉此处才可以读取数据
	minios, _ := connectMinio("127.0.0.1:9000", "admin", "admin123456")
	n, err := minios.PutObject("db5", "haha.png", file, fileStat, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	fmt.Println(n, err)
}
