package internal_minio

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

func GeneratePresignedURL(bucketName, objectName string, expiry time.Duration, queryParams map[string]string) (string, error) {
	ctx := context.Background()

	reqParams := make(url.Values)
	for key, value := range queryParams {
		reqParams.Set(key, value)
	}

	presignedURL, err := MinioClient.PresignedGetObject(ctx, bucketName, objectName, expiry, reqParams)
	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}

func UploadFile(bucketName, objectName, filePath, contentType string) (*minio.UploadInfo, error) {
	ctx := context.Background()

	exists, err := MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			return nil, err
		}
	}

	info, err := MinioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func DownloadFile(ctx context.Context, bucketName, objectName string) (io.Reader, int64, string, error) {
	object, err := MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, 0, "", err
	}

	objectInfo, err := object.Stat()
	if err != nil {
		return nil, 0, "", err
	}

	return object, objectInfo.Size, objectInfo.ContentType, nil
}

func DeleteFile(bucketName, objectName string) error {
	ctx := context.Background()

	err := MinioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

func ListFiles(bucketName string) ([]string, error) {
	ctx := context.Background()
	var files []string

	objectCh := MinioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{Recursive: true})
	for object := range objectCh {
		if object.Err != nil {
			return nil, object.Err
		}
		files = append(files, object.Key)
	}

	return files, nil
}

func DeleteBucket(bucketName string) error {
	ctx := context.Background()

	err := MinioClient.RemoveBucket(ctx, bucketName)
	if err != nil {
		return err
	}

	return nil
}

func Exist(bucketName, objectName string) bool {
	ctx := context.Background()

	_, err := MinioClient.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false
		}
		return false
	}

	return true
}
