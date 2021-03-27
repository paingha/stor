package s3

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3ConnectionDetails struct {
	bucket string
	region string
}

func (s *S3ConnectionDetails) UploadExportToS3(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file %s %s", filename, err)
	}
	defer file.Close()
	session := session.New(&aws.Config{
		Region: aws.String(s.region),
	})
	svc := s3manager.NewUploader(session)
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("File Upload error %s %s", filename, err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", filename, result.Location)
}
