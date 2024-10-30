package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type fsBucket struct {
	fsb *gridfs.Bucket
}

type FSBucketClient interface {
	OpenUploadStream(ctx context.Context, fileId string, chuckSize int32, metaData interface{}) (*gridfs.UploadStream, error)
	OpenDownloadStreamByName(ctx context.Context, filename string) (*gridfs.DownloadStream, error)
	OpenDownloadStreamByFileID(ctx context.Context, fileID interface{}) (*gridfs.DownloadStream, error)
	Delete(ctx context.Context, fileID interface{}) error
}

func (f *fsBucket) OpenUploadStream(ctx context.Context, fileId string, chuckSize int32, metaData interface{}) (*gridfs.UploadStream, error) {
	return f.fsb.OpenUploadStream(fileId, options.GridFSUpload().SetChunkSizeBytes(chuckSize).SetMetadata(metaData))
}

func (f *fsBucket) OpenDownloadStreamByName(ctx context.Context, filename string) (*gridfs.DownloadStream, error) {
	return f.fsb.OpenDownloadStreamByName(filename)
}

func (f *fsBucket) OpenDownloadStreamByFileID(ctx context.Context, fileID interface{}) (*gridfs.DownloadStream, error) {
	return f.fsb.OpenDownloadStream(fileID)
}

func (f *fsBucket) Delete(ctx context.Context, fileID interface{}) error {
	return f.fsb.Delete(fileID)
}

func NewFSBucket(mgConnector *mongo.Client, databaseName string) (FSBucketClient, error) {
	db := mgConnector.Database(databaseName)
	fsb, err := gridfs.NewBucket(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create gridfs bucket: %w", err)
	}
	return &fsBucket{fsb: fsb}, nil
}
