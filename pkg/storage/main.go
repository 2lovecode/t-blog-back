package storage

import (
	"t-blog-back/pkg/logging"
	"t-blog-back/pkg/setting"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// TStorage 存储
var TStorage *minio.Client

// SetUp 初始化
func SetUp() {
	endpoint := setting.StorageCfg.Endpoint
	accessKeyID := setting.StorageCfg.AccessKey
	secretAccessKey := setting.StorageCfg.SecretKey

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logging.Tlog.Fatalln(err)
	}

	// minioClient.SetBucketPolicy(context.TODO(), setting.StorageCfg.BucketName, "public")

	logging.Tlog.Printf("%#v\n", minioClient) // minioClient初使化成功
}
