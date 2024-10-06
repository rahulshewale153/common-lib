package mongodb

import (
	"bytes"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnector struct {
	client   *mongo.Client
	fsBucket *gridfs.Bucket
}

type MongoDBConfig struct {
	URI      string
	Database string
}

//URI format: mongodb://username:password@host:port

func NewMongoDBConnector(ctx context.Context, cfg *MongoDBConfig) (*MongoDBConnector, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	fmt.Println("Connected to MongoDB!")
	database := client.Database(cfg.Database)
	fsBucket, err := gridfs.NewBucket(database)
	if err != nil {
		return nil, fmt.Errorf("failed to create gridfs bucket: %w", err)
	}

	return &MongoDBConnector{
		client:   client,
		fsBucket: fsBucket,
	}, nil
}

// method to upload file to mongodb
func (conn *MongoDBConnector) Upload(ctx context.Context, fileId string, file []byte, chuckSize int32, metaData interface{}) (interface{}, error) {
	// Create an upload stream with GridFS
	uploadStream, err := conn.fsBucket.OpenUploadStream(fileId, options.GridFSUpload().SetChunkSizeBytes(chuckSize).SetMetadata(metaData))
	if err != nil {
		return "", fmt.Errorf("failed to open upload stream: %w", err)
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(file)
	if err != nil {
		return "", fmt.Errorf("failed to write file data: %w", err)
	}

	// Return the ObjectId of the uploaded file
	return uploadStream.FileID, nil
}

// method to download file from mongodb
func (conn *MongoDBConnector) Download(ctx context.Context, fileId string) (string, error) {
	var buf bytes.Buffer
	downloadStream, err := conn.fsBucket.OpenDownloadStreamByName(fileId)
	if err != nil {
		return "", fmt.Errorf("failed to open download stream: %w", err)
	}
	defer downloadStream.Close()

	// Read the file content into the buffer
	_, err = buf.ReadFrom(downloadStream)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// method to delete file from mongodb
func (conn *MongoDBConnector) Delete(ctx context.Context, fileId interface{}) error {
	err := conn.fsBucket.Delete(fileId)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w %s", err, fileId)
	}
	return nil
}

func (conn *MongoDBConnector) Close(ctx context.Context) error {
	err := conn.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("failed to disconnect from mongodb: %w", err)
	}
	return nil
}
