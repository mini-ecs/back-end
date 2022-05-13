package service

import (
	"encoding/json"
	"github.com/mini-ecs/back-end/internal/model"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/assert"
	"log"
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
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up

	minio.NewPostPolicy().
}
