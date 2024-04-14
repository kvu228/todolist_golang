package components

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"time"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) error
	GetName() string
}

type s3Provider struct {
	id         string
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func (p *s3Provider) ID() string { return p.id }

func NewAWSS3Provider(id string) *s3Provider {
	return &s3Provider{id: id}
}

func (p *s3Provider) SaveFileUploaded(ctx context.Context, data []byte, dst string) error {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	_, err := s3.New(p.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(p.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *s3Provider) GetUploadPresignedURL(ctx context.Context) string {
	req, _ := s3.New(p.session).PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(p.bucketName),
		Key:    aws.String(fmt.Sprintf("img/%d", time.Now().UnixNano())),
		ACL:    aws.String("private"),
	})
	//
	url, _ := req.Presign(time.Second * 60)

	return url
}

func (p *s3Provider) GetDomain() string { return p.domain }
func (*s3Provider) GetName() string     { return "aws_s3" }
