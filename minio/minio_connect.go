package minio

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"os"
)

// connect minio
func connectMinio(ipaddr, accessKeyID, secretAccessKey string) (*minio.Client, error) {
	ssl := false
	// init minio client object.
	//minioClient, err := minio.New("127.0.0.1:9000", "admin", "admin123456", ssl)

	minioClient, err := minio.New(ipaddr, accessKeyID, secretAccessKey, ssl)
	if err != nil {
		fmt.Println(err)
		panic(err)
		return nil, nil
	}
	return minioClient, err
}

// create Bucket

func creatBucket(minioClient *minio.Client, bucketName string) error {
	found, err := minioClient.BucketExists(bucketName)
	if err != nil {
		return err
	}
	if found {
		return nil
	}
	err = minioClient.MakeBucket(bucketName, "cn-north-1")
	return err
}

// put Object

func PutObject(minioClient *minio.Client, bucketName string, objectName string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return err
	}

	n, err := minioClient.PutObject(bucketName, objectName, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return err
}
