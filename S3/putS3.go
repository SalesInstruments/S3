package S3

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (c ConfigS3) PutFile(client *s3.Client, file *multipart.FileHeader) {

	// Задаем параметры для загрузки
	src, err := file.Open()
	if err != nil {
		log.Fatalf("Failed to open file, %v", err)
	}
	defer src.Close()

	// Задаем параметры для загрузки
	key := file.Filename

	// Загружаем файл в S3
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
		Body:   src,                          // Передаем тело файла
		ACL:    types.ObjectCannedACLPrivate, // Публичный или приватный доступ
	})
	if err != nil {
		log.Fatalf("Failed to upload file, %v", err)
	}

	fmt.Printf("File uploaded successfully to %s/%s\n", c.Bucket, key)
}
