package s3

import (
	"fmt"
	"os"
	"thebrag/requests"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadObject(bucket string, filePath string, fileName string, sess *session.Session, awsConfig requests.AWSConfig) error {

	// Open file to upload
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Unable to open file %v", err)
		return err
	}
	defer file.Close()

	// Upload to s3
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		fmt.Printf("failed to upload object, %v\n", err)
		return err
	}

	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	return nil
}

func UploadFile(filePath string, awsConfig requests.AWSConfig) error {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(awsConfig.Region),
		Credentials: credentials.NewStaticCredentials(
			awsConfig.AccessKeyID,
			awsConfig.AccessKeySecret,
			"",
		),
		Endpoint: &awsConfig.BaseURL,
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", filePath, err)
	}

	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(filePath),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	return nil
}
