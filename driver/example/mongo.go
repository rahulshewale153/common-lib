package main

import (
	"context"
	"fmt"

	"github.com/rahulshewale153/common-lib/driver/mongodb"
)

func main() {
	mongoConfig := &mongodb.MongoDBConfig{
		Host:     "localhost",
		Port:     "27017",
		Username: "root", // MongoDB Username
		Password: "root", // MongoDB Password
		Database: "admin",
	}

	conn, err := mongodb.NewMongoDBConnector(context.Background(), mongoConfig)
	if err != nil {
		fmt.Println("Error while connecting to MongoDB!!!")
		return
	}

	fmt.Println("Connected to MongoDB!!!")

	// Upload file to MongoDB
	fileId := "file1"
	file := []byte("Hello, World!")
	chunkSize := int32(261120)
	metaData := map[string]string{
		"filename": "file1",
		"mimetype": "text/plain",
	}

	objectId, err := conn.Upload(context.Background(), fileId, file, chunkSize, metaData)
	if err != nil {
		fmt.Println("Error while uploading file to MongoDB!!!")
		return
	}

	fmt.Println("File uploaded successfully!!!")

	//Download file from MongoDB
	downloadedFile, err := conn.Download(context.Background(), fileId)
	if err != nil {
		fmt.Println("Error while downloading file from MongoDB!!!")
		return
	}

	fmt.Println("File downloaded successfully!!!")
	fmt.Println("File Data: ", string(downloadedFile))

	// Delete file from MongoDB
	err = conn.Delete(context.Background(), objectId)
	if err != nil {
		fmt.Println("Error while deleting file from MongoDB!!!", err)
		return
	}

	fmt.Println("File deleted successfully!!!")

}
