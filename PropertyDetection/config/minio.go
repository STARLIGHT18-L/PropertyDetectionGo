package config

import (
	"fmt"
	"github.com/minio/minio-go"
)

func MinioBucketInit(minioClient *minio.Client) {
	exists, err := minioClient.BucketExists(Boot.Config.Minio.Bucket)
	if err != nil {
		panic(err)
	}
	if !exists {
		// 创建桶
		err = minioClient.MakeBucket(Boot.Config.Minio.Bucket, "")
		if err != nil {
			panic(err)
		}
		// 定义桶策略
		policy := fmt.Sprintf(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Principal": {"AWS": ["*"]},
					"Action": ["s3:GetObject"],
					"Resource": ["arn:aws:s3:::%s/*"]
				}
			]
		}`, Boot.Config.Minio.Bucket)
		// 设置桶策略
		err = minioClient.SetBucketPolicy(Boot.Config.Minio.Bucket, policy)
		if err != nil {
			panic(err)
		}
		fmt.Println("创建存储桶成功==>", Boot.Config.Minio.Bucket)
	} else {
		fmt.Println("存储桶已存在==>", Boot.Config.Minio.Bucket)
	}
}
