package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createCognitoUserPool(ctx *pulumi.Context) (*cognito.UserPool, error) {
	userPool, err := cognito.NewUserPool(ctx, "ragnarUserPool", &cognito.UserPoolArgs{
		AutoVerifiedAttributes: pulumi.StringArray{
			pulumi.String("email"),
		},
		UsernameAttributes: pulumi.StringArray{
			pulumi.String("email"),
		},
		AdminCreateUserConfig: &cognito.UserPoolAdminCreateUserConfigArgs{
			AllowAdminCreateUserOnly: pulumi.Bool(true),
		},
	})
	errorHandler(err)

	return userPool, err
}
func createCognitoUserPoolClient(ctx *pulumi.Context, userPool *cognito.UserPool) (*cognito.UserPoolClient, error) {
	userPoolClient, err := cognito.NewUserPoolClient(ctx, "ragnarUserPoolClient", &cognito.UserPoolClientArgs{
		UserPoolId: userPool.ID(),
		ExplicitAuthFlows: pulumi.StringArray{
			pulumi.String("ALLOW_USER_PASSWORD_AUTH"),
			pulumi.String("ALLOW_REFRESH_TOKEN_AUTH"),
		},
		GenerateSecret: pulumi.Bool(true),
	})
	errorHandler(err)

	return userPoolClient, nil
}
func createCogntoDomain(ctx *pulumi.Context, userPool *cognito.UserPool) (*cognito.UserPoolDomain, error) {
	domain, err := cognito.NewUserPoolDomain(ctx, "RagarUserPoolDomain", &cognito.UserPoolDomainArgs{
		Domain:     pulumi.String("ragnar-cli"),
		UserPoolId: userPool.ID(),
	})
	errorHandler(err)

	_, err = acm.NewCertificate(ctx, "ragnarCliCert", &acm.CertificateArgs{
		DomainName:       pulumi.String("ragnar.cli.the0x.dev"),
		ValidationMethod: pulumi.String("DNS"),
	})
	errorHandler(err)

	_, err = cognito.NewUser(ctx, "foxhound", &cognito.UserArgs{
		UserPoolId: userPool.ID(),
		Username:   pulumi.String("foxhound@the0x.dev"),
		Attributes: pulumi.StringMap{
			"email":          pulumi.String("foxhound@the0x.dev"),
			"email_verified": pulumi.String("true"),
		},
	})
	errorHandler(err)

	ctx.Export("cognitoDomain", domain.Domain)

	return domain, nil
}
