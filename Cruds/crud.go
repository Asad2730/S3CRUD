package cruds

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// if want to store in bucket then replace Object with Bucket in every fuction

func Create_Update(bucketName string, objectKey string, svc *s3.S3) {

	filePath := "/path/to/your_file.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("error opening file:", err.Error())
		os.Exit(1)
	}

	defer file.Close()

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	}

	_, err = svc.PutObject(input)

	if err != nil {
		fmt.Println("Error Uploading Data to s3", err.Error())
	}

	fmt.Println("Data inserted successfully!")

}

func Read(bucketName string, objectKey string, svc *s3.S3) {

	downloadInput := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	result, err := svc.GetObject(downloadInput)

	if err != nil {
		fmt.Println("Error retrieving data from s3", err.Error())
		os.Exit(1)
	}

	filePath := "/path_to_your_new_file_were_you_want to store retrived data/file.txt"

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file", err.Error())
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.ReadFrom(result.Body)
	if err != nil {
		fmt.Println("Error writing data to file", err.Error())
		os.Exit(1)
	}

	fmt.Println("Data retrieved (downloaded) successfully.")
}

func Delete(bucketName string, objectKey string, svc *s3.S3) {

	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	_, err := svc.DeleteObject(deleteInput)

	if err != nil {
		fmt.Println("Error deleting data from s3:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Data deleted successfully")
}
