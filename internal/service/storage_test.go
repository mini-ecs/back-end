package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"
	"log"
	"os/exec"
	"testing"
)

const conf = `{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:ListBucket"
            ],
            "Resource": [
                "arn:aws:s3:::*"
            ],
            "Condition": {
                "StringLike": {
                    "s3:prefix": [
                        "${aws:username}/*"
                    ]
                }
            }
        },
        {
            "Effect": "Allow",
            "Action": [
                "s3:*"
            ],
            "Resource": [
                "arn:aws:s3:::${aws:username}/*"
            ]
        }
    ]
}`

func TestUnmarshal(t *testing.T) {
	con := model.UserAccessConf{}
	err := json.Unmarshal([]byte(conf), &con)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, TemplateUserConf, con)
	con1, err := json.Marshal(con)
	if err != nil {
		t.Error(err)
	}
	con2, err := json.Marshal(TemplateUserConf)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, con1, con2)
}

func TestConstValue(t *testing.T) {
	assert.Equal(t, "arn:aws:s3:::${aws:username}/*", CommonUserFolderAndChildren)
	assert.Equal(t, "${aws:username}/*", PermissionByUserName+AllFolder)
	assert.Equal(t, "arn:aws:s3:::*", Prefix+"*")
}

func TestConnnect(t *testing.T) {
	endpoint := "localhost:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
	err = minioClient.MakeBucket(context.Background(), "testbucket", minio.MakeBucketOptions{})
	if err != nil {
		log.Fatalln(err)
	}
}

func Test_MakeBucket(t *testing.T) {
	cmd := exec.Command("mc", "mb", "bucket")
	fmt.Printf("%v %v, %v", cmd.Path, cmd.Args, cmd.String())
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
