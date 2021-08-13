package src

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFile(file Target, config Configuration) (int, error) {
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWS.Region),
	})
	if err != nil {
		return 0, err
	}

	fp, err := os.Open(file.LocalFilename)
	if err != nil {
		return 0, fmt.Errorf("Unable to open file when trying to upload: %s", err)
	}
	defer (func() {
		err := fp.Close()
		if err != nil {
			fmt.Println("Failed to close file: %s", err)
		}
	})()

	fileInfo, _ := fp.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	byteCount, err := fp.Read(buffer)

	if err != nil {
		return byteCount, fmt.Errorf("could not open file: %s", err)
	}

	s3Key, _ := file.GetS3ObjectKey(config)

	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:        &config.BucketName,
		Key:           &s3Key,
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(size),
	})

	return byteCount, err
}
