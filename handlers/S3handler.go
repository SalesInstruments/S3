package handlers

import (
	"S3/S3"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var s3config = S3.ConfigS3{
	URL:        "https://s3.timeweb.cloud",
	Key:        "97DW7TQ210XFB53RLJ93",
	Secret_key: "tU0XVxOf5LbNDiBKMlC2WjeUv2hcCLWa3sApobeQ",
	Bucket:     "b6c3cbcd-db7de6ad-19e4-40d3-a077-697251c8babf",
	Region:     "ru-1",
}

func UploadFile(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Ошибка при получении файла")
	}

	factory := S3.S3ClientFactory{}
	s3Client, err := factory.CreateS3Client(s3config)

	if err != nil {
		log.Fatalf("Error creating S3 client: %v", err)
	}

	s3config.PutFile(s3Client, file)

	return nil
}

func DeleteFile(c *fiber.Ctx) error {

	fileName := c.Params("fileName")
	fmt.Println(fileName + "kljlkjlk")

	factory := S3.S3ClientFactory{}
	s3Client, err := factory.CreateS3Client(s3config)

	if err != nil {
		log.Fatalf("Error creating S3 client: %v", err)
	}

	s3config.DelFile(s3Client, fileName)

	return nil
}
