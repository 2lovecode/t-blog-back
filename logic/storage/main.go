package storage

import (
	"context"
	"t-blog-back/pkg/setting"
	ss "t-blog-back/pkg/storage"

	"github.com/minio/minio-go/v7"
)

type StorageLogic struct {
	StorageClient *minio.Client
	BucketName    string
}

func NewStorageLogic(ctx context.Context) StorageLogic {
	storageLogic := StorageLogic{}
	storageLogic.StorageClient = ss.TStorage
	storageLogic.BucketName = setting.StorageCfg.BucketName

	storageLogic.StorageClient.SetBucketPolicy(ctx, storageLogic.BucketName, "public")
	return storageLogic
}

func Put(ctx context.Context) {

}

func Get(ctx context.Context) {

}
