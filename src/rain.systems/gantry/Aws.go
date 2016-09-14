package gantry

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/aws"

	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func EcrLogin(conf Config) *string {
	fmt.Println("Auth")

	var cred *credentials.Credentials
	if(len(os.Getenv("AWS_ACCESS_KEY_ID")) > 0 && len(os.Getenv("AWS_ACCESS_KEY")) > 0) {
		cred = credentials.NewEnvCredentials()
	}
	if(len(conf.Aws.AccessId) > 0 && len(conf.Aws.SecretKey) > 0) {
		cred = credentials.NewStaticCredentials(conf.Aws.AccessId, conf.Aws.SecretKey, "")
	}

	awsConf := &aws.Config{
		Credentials: cred,
		Region: &conf.Aws.Region,
	}

	svc := ecr.New(session.New(awsConf))

	params := &ecr.GetAuthorizationTokenInput{
		RegistryIds: []*string{
			aws.String(conf.Docker.Tag[0:12]), // Required
			// More values...
		},
	}
	resp, err := svc.GetAuthorizationToken(params)
	if err != nil {
		log.Fatalln(err);
		os.Exit(1);
	}

	fmt.Println(conf.Docker.Tag)
	fmt.Println(resp.GoString())

	return resp.AuthorizationData[0].AuthorizationToken
}
