package config

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func InitS3() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1")) // Adjust region
	if err != nil {
		log.Fatalf("unable to load AWS config, %v", err)
	}
	s3Client = s3.NewFromConfig(cfg)
}

func UploadFile(file multipart.File, filename string) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return "", err
	}

	bucket := os.Getenv("AWS_BUCKET")
	key := "uploads/" + filename

	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   buffer,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key), nil
}
