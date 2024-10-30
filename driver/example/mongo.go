package main

import (
	"bytes"
	"context"
	"fmt"

	"github.com/rahulshewale153/common-lib/driver/mongodb"
)

func main() {
	mongoConfig := &mongodb.MongoDBConfig{
		Host:     "localhost",
		Port:     27017,
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
	fsBucket, err := mongodb.NewFSBucket(conn.Client, "admin")
	if err != nil {
		fmt.Println("Error while creating file bucket!!!")
		return
	}
	// Upload file to MongoDB
	fileName := "file1"
	file := []byte("Hello, World!")
	chunkSize := int32(261120)
	metaData := map[string]string{
		"filename": fileName,
		"mimetype": "text/plain",
	}

	openUploadStream, err := fsBucket.OpenUploadStream(context.Background(), fileName, chunkSize, metaData)
	if err != nil {
		fmt.Println("Error while opening upload stream!!!")
		return
	}

	_, err = openUploadStream.Write(file)
	if err != nil {
		fmt.Errorf("failed to write file data: %w", err)
		return
	}

	//before downloading and deleting file we need to close the upload
	openUploadStream.Close()

	objectId := openUploadStream.FileID
	fmt.Println("File uploaded successfully!!!")

	//Download file from MongoDB by file name
	openDownloadStream, err := fsBucket.OpenDownloadStreamByName(context.Background(), fileName)
	if err != nil {
		fmt.Println("Error while opening download stream!!!")
		return
	}

	var buf bytes.Buffer
	// Read the file content into the buffer
	_, err = buf.ReadFrom(openDownloadStream)
	if err != nil {
		fmt.Println("Error while reading download stream!!!")
		return
	}

	openDownloadStream.Close()

	fmt.Println("File downloaded successfully by file Name!!!")
	fmt.Println("File Data: ", buf.String())

	//Download file from MongoDB by fileid

	openDownloadStream, err = fsBucket.OpenDownloadStreamByFileID(context.Background(), objectId)
	if err != nil {
		fmt.Println("Error while opening download stream!!!", err)
		return
	}

	buf.Reset()
	// Read the file content into the buffer
	_, err = buf.ReadFrom(openDownloadStream)
	if err != nil {
		fmt.Println("Error while reading download stream!!!")
		return
	}

	defer openDownloadStream.Close()

	fmt.Println("File downloaded successfully by fileid!!!")
	fmt.Println("File Data: ", buf.String())

	// Delete file from MongoDB
	err = fsBucket.Delete(context.Background(), objectId)
	if err != nil {
		fmt.Println("Error while deleting file from MongoDB!!!", err)
		return
	}

	fmt.Println("File deleted successfully!!!")

}
