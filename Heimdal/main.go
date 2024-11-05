package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := createBucket(ctx)
		errorHandler(err)

		role, err := createIAMRole(ctx, bucket)
		errorHandler(err)

		_, err = createLambdaFunc(ctx, role, bucket)
		errorHandler(err)

		cert, err := createACMCertificates(ctx)
		errorHandler(err)

		_, err = createApiGateway(ctx, cert)
		errorHandler(err)

		userPool, err := createCognitoUserPool(ctx)
		errorHandler(err)

		_, err = createCognitoUserPoolClient(ctx, userPool)
		errorHandler(err)

		_, err = createCogntoDomain(ctx, userPool)
		errorHandler(err)

		err = createBillingAlarm(ctx)
		errorHandler(err)

		return nil
	})
}
