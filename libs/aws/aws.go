package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type AWSLibrary struct{}

// AWSLibraryHandler ...
func AWSLibraryHandler() *AWSLibrary {
	return &AWSLibrary{}
}

// AWSLibraryInterface ...
type AWSLibraryInterface interface {
	GetSystemManagerStore() *ssm.SSM
}

func sessions() *aws.Config {
	creds := credentials.NewStaticCredentials(
		os.Getenv("AWS_ACCESS_KEY"),
		os.Getenv("AWS_ACCESS_SECRET"), "")
	creds.Get()
	cfgAws := aws.NewConfig().WithRegion(os.Getenv("AWS_ACCESS_AREA")).WithCredentials(creds)
	return cfgAws
}

// GetSystemManagerStore ...
func (aws *AWSLibrary) GetSystemManagerStore() *ssm.SSM {
	cfg := sessions()
	return ssm.New(session.New(), cfg)
}
