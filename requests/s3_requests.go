package requests

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	UploadTimeout   int
	BaseURL         string
}
