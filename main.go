package main

import (
	"fmt"
	"os"

	cruds "github.com/Asad2730/GoAWS/Cruds"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("Your_Region"),
	})

	if err != nil {
		fmt.Println("Error creating aws session", err.Error())
		os.Exit(1)
	}

	svc := s3.New(sess)

	bucketName := "your_bucket_name"
	objectKey := "your_object_key"

	//to update just send same  key it will update pre and insert new one there
	cruds.Create_Update(bucketName, objectKey, svc)
	cruds.Read(bucketName, objectKey, svc)
	cruds.Delete(bucketName, objectKey, svc)
}
