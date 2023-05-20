package main

import (
	"context"
	"fmt"
	"github.com/anisurrahman75/go-cloud-blob/blob"
	"github.com/anisurrahman75/go-cloud-blob/model"
)

type Blob interface {
	Upload(ctx context.Context, filepath string, data []byte) error
	Get(ctx context.Context, filepath string) ([]byte, error)
	List(ctx context.Context, dir string) ([][]byte, error)
	Delete(ctx context.Context, filepath string, isDir bool) error
	Exists(ctx context.Context, filepath string) (bool, error)
}

func NewTestS3() (Blob, error) {
	bs := &model.BackupStorage{
		Storage: model.Storage{
			Provider: model.ProviderS3,
			S3: &model.S3{
				Bucket:   "",
				Region:   "",
				Endpoint: "",
				Prefix:   "",
			},
		},
	}
	return blob.NewBlob(bs)
}
func main() {
	b, err := NewTestS3()
	if err != nil {
		fmt.Print(err.Error())
	}
	_ = b
}
