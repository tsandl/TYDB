// 实现minio的一些基本操作
package miomanager

import (
	"context"
	"fmt"
	_minio "github.com/minio/minio-go/v6"
	"io"
)

type Client struct {
	mClient _minio.Client
}

func NewMinIO(ipaddr, accessName, accessPasswd string) *Client {
	ssl := false // close http
	// 初使化minio client对象。
	//minioClient, err := _minio.New(ipaddr, access_Name, access_Passwd, ssl)
	minioClient, err := _minio.New(ipaddr, accessName, accessPasswd, ssl)
	if err != nil {
		panic(err)
	}
	return &Client{*minioClient}
}

func (cl *Client) makeBucket(bucketName string) error {
	found, err := cl.mClient.BucketExists(bucketName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if found {
		fmt.Println("Bucket found")
		return err
	}

	err = cl.mClient.MakeBucket(bucketName, "cn-north-1")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully created mybucket.")
	return err
}

func (cl *Client) removeBucket(bucketName string) error {
	f, _ := cl.mClient.BucketExists(bucketName)
	if f != true { // if bucket is not exist, return directly
		fmt.Println("Bucket is not exist.....")
		return nil
	}
	// if bucket exist, remove bucket
	err := cl.mClient.RemoveBucket(bucketName)
	if err != nil {
		panic(err)
	}
	return nil
}

func (cl *Client) PutObject(bucketName, objectName string, read io.Reader, objectSize int64) (l int64, err error) {
	f, _ := cl.mClient.BucketExists(bucketName)
	if f != true { // if bucket is not exist, return directly
		fmt.Println("Bucket is not exist.....")
		return 0, nil
	}
	// 此处的putObjectOptions操作默认为“application/octet-stream” ，application/octet-stream 默认为未知文件
	l, err = cl.mClient.PutObject(bucketName, objectName, read, objectSize, _minio.PutObjectOptions{ContentType: "application/octet-stream"})
	return l, err
}
func (cl *Client) removeObject(bucketName, objectName string) error {
	err := cl.mClient.RemoveObject(bucketName, objectName)
	return err
}
func (cl *Client) removeObjects(bucketName string, objectsCh chan string) {

	for rErr := range cl.mClient.RemoveObjects(bucketName, objectsCh) {
		fmt.Println("Error detected during deletion: ", rErr)
	}
}

func (cl *Client) fPutObject(bucketName, objectName, filePath string) (length int64, err error) {
	length, err = cl.mClient.FPutObject(bucketName, objectName, filePath, _minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println("Error happen when putFile: ", err)
	}
	return length, err
}

// 与fPutObject很像，不过可以取消上传
func (cl *Client) fPutObjectWithContext(ctx context.Context, bucketName, objectName, filePath string) (length int64, err error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	//defer cancel()
	// ctx 需要用户在自己的从层面写
	length, err = cl.mClient.FPutObjectWithContext(ctx, bucketName, objectName, filePath, _minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully uploaded bytes: ", length)
	return length, err
}

func (cl *Client) getObject(bucketName, objectName string) (*_minio.Object, error) {
	object, err := cl.mClient.GetObject(bucketName, objectName, _minio.GetObjectOptions{})
	return object, err
}

func (cl *Client) fGetObject(bucketName, objectName, filePath string) error {

	err := cl.mClient.FGetObject(bucketName, objectName, filePath, _minio.GetObjectOptions{})
	return err
}
func (cl *Client) getObjectWithContext(ctx context.Context, bucketName, objectName string) (*_minio.Object, error) {
	//ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	//defer cancel()
	object, err := cl.mClient.GetObjectWithContext(ctx, bucketName, objectName, _minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return object, err
}

func (cl *Client) fGetObjectWithContext(ctx context.Context, bucketName, objectName, filePath string) error {
	err := cl.mClient.FGetObjectWithContext(ctx, bucketName, objectName, filePath, _minio.GetObjectOptions{})
	return err
}
