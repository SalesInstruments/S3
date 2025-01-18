package S3

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c ConfigS3) DelFile(client *s3.Client, fileName string) {
	// Prepare the delete object input
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(c.Bucket), // s3 bucket name
		Key:    aws.String(fileName), // file name
	}

	// Delete the object
	_, err := client.DeleteObject(context.TODO(), input)
	if err != nil {
		log.Fatalf("не удалось удалить файл: %v", err)
	}
}
