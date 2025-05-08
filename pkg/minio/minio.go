package internal_minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"webapi/config"
)

var (
	MinioClient *minio.Client
)

func Setup() error {
	var minioClient *minio.Client
	ctx := context.Background()

	conf := config.GetConfig()

	if conf.Minio.Enable {
		minioClient, _ = minio.New(conf.Minio.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(conf.Minio.AccessKeyID, conf.Minio.AccessKeySecret, ""),
			Secure: conf.Minio.UseSSL,
		})

		if _, err := minioClient.ListBuckets(ctx); err != nil {
			return err
		}
	}

	MinioClient = minioClient

	return nil
}

func IsAlive() bool {
	ctx := context.Background()

	if _, err := MinioClient.ListBuckets(ctx); err != nil {
		return false
	}

	return true
}

func GetMinio() *minio.Client {
	return MinioClient
}
