package gantry

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/aws"

)

func ecrLogin() {

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := ecr.New(sess)

	params := &ecr.GetAuthorizationTokenInput{
		RegistryIds: []*string{
			aws.String("RegistryId"), // Required
			// More values...
		},
	}
	resp, err := svc.GetAuthorizationToken(params)

	fmt.Println(resp.GoString())

	dockerLogin()
}
