package components

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"time"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) error
	GetName() string
	GetDomain() string
}

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewAWSS3Provider(bucketName, region, apiKey, secret, domain string) UploadProvider {
	newSession, _ := session.NewSession(&aws.Config{Region: aws.String(region), Credentials: credentials.NewStaticCredentials(apiKey, secret, "")})
	return &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
		session:    newSession,
	}
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
