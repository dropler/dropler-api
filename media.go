package main

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
)

var (
	BUCKET_NAME = "droppio"

	S3 *s3.S3
)

func init() {
	creds := credentials.NewEnvCredentials()
	_, err := creds.Get()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	config := &aws.Config{
		Region:           aws.String("us-east-1"),
		Endpoint:         aws.String("s3.amazonaws.com"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      creds,
	}

	S3 = s3.New(config)

}

func UploadMedia(c *gin.Context) {

	// Get authorized user
	u := c.MustGet("user").(*User)

	// the FormFile function takes in the POST input id file
	file, _, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to get file from post body", "err": err})
		c.AbortWithStatus(400)
		return
	}
	defer file.Close()

	// Generate ID for Media
	mediaID := uuid.New()

	objectKey := fmt.Sprintf("%x/%s", u.ID, mediaID)

	params := &s3.PutObjectInput{
		Bucket: aws.String(BUCKET_NAME),
		Key:    aws.String(objectKey),
		Body:   file,
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
	}

	_, err = S3.PutObject(params)

	if err != nil {
		c.JSON(400, gin.H{"message": "Failed to upload image", "error": err})
		c.AbortWithStatus(400)
		return
	}

	c.JSON(200, gin.H{"message": "SUCCESS", "id": mediaID})
}

func GetMedia(c *gin.Context) {
	// Get authorized user
	u := c.MustGet("user").(*User)
	mediaID := c.Param("id")

	objectKey := fmt.Sprintf("%x/%s", u.ID, mediaID)

	params := &s3.GetObjectInput{
		Bucket: aws.String(BUCKET_NAME), // Required
		Key:    aws.String(objectKey),   // Required
	}

	resp, err := S3.GetObject(params)

	if err != nil {
		c.JSON(404, gin.H{"message": "NOT FOUND", "error": err})
		c.AbortWithStatus(404)
		return
	}

	// Copy the media buffer into the response writer
	io.Copy(c.Writer, resp.Body)
}
