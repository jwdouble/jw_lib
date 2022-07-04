package minix

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"log"
	"testing"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Test_minio(t *testing.T) {
	endpoint := "150.158.7.96:9000"
	accessKeyID := "admin"
	secretAccessKey := "admin123"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//log.Printf("%#v\n", minioClient) // minioClient is now set up
	r, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(r)
}
