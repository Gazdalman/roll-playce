package config

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"


	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func InitS3() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-2")) // Adjust region
	if err != nil {
		log.Fatalf("unable to load AWS config, %v", err)
	}
	s3Client = s3.NewFromConfig(cfg)
}

func getUniqueFilename(filename string) string {
	// Extract the file extension (if any)
	ext := strings.ToLower(filepath.Ext(filename))
	if ext != "" {
		ext = ext[1:] // Remove the leading dot
	}
	// Generate a unique UUID
	uniqueFilename := uuid.New().String()
	return fmt.Sprintf("%s.%s", uniqueFilename, ext)
}

func UploadFile(file multipart.File, filename string) (string, error) {
	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		return "", err
	}

	bucket := os.Getenv("AWS_BUCKET")

	// Generate a unique filename
	filename = getUniqueFilename(filename)

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
