package minio

import (
	"context"
	_minio "github.com/minio/minio-go/v6"
	"io"
	"time"
)

type BucketInfo struct {
	bucketName         string
	bucketCreationDate time.Time
}

type ObjectInfo struct {
	objectInfoKey          string
	objectInfoSize         int64
	objectInfoETag         string
	objectInfoLastModified time.Time
}

type MINIO interface {
	makeBucket(bucketName string) error
	//listBuckets() ([]BucketInfo, error)
	removeBucket(bucketName string) error
	//listObjects(bucketName, prefix string, recursive bool, doneCh chan struct{}) <-chan ObjectInfo
	putObject(bucketName, objectName string, reader io.Reader, objectSize int64) (l int64, err error)
	removeObject(bucketName, objectName string) error
	removeObjects(bucketName string, objectsCh chan string) // delete objects
	fPutObject(bucketName, objectName, filePath string) (length int64, err error)
	fPutObjectWithContext(ctx context.Context, bucketName, objectName, filePath string) (length int64, err error)
	//getObject(bucketName, objectName string) (* _minio.GetObjectOptions, error)
	getObject(bucketName, objectName string) (*_minio.Object, error)
	fGetObject(bucketName, objectName, filePath string) error
	getObjectWithContext(ctx context.Context, bucketName, objectName string) (*_minio.Object, error)
	fGetObjectWithContext(ctx context.Context, bucketName, objectName, filePath string) error
}
