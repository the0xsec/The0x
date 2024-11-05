package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createApiGateway(ctx *pulumi.Context, cert *acm.Certificate) (*apigatewayv2.Api, error) {
	api, err := apigatewayv2.NewApi(ctx, "ragnarApi", &apigatewayv2.ApiArgs{
		Name:         pulumi.String("ragnar-api"),
		ProtocolType: pulumi.String("HTTP"),
	})
	errorHandler(err)

	stage, err := apigatewayv2.NewStage(ctx, "ragnarStaging", &apigatewayv2.StageArgs{
		ApiId:      api.ID(),
		Name:       pulumi.String("ragnar-api-stage"),
		AutoDeploy: pulumi.Bool(true),
	})
	errorHandler(err)

	domain, err := apigatewayv2.NewDomainName(ctx, "ragnarApiDomain", &apigatewayv2.DomainNameArgs{
		DomainName: pulumi.String("ragnar.the0x.dev"),
		DomainNameConfiguration: &apigatewayv2.DomainNameDomainNameConfigurationArgs{
			CertificateArn: cert.Arn,
			EndpointType:   pulumi.String("REGIONAL"),
			SecurityPolicy: pulumi.String("TLS_1_2"),
		},
	})
	errorHandler(err)

	_, err = apigatewayv2.NewApiMapping(ctx, "ragnarApiMapping", &apigatewayv2.ApiMappingArgs{
		ApiId:      api.ID(),
		DomainName: domain.ID(),
		Stage:      stage.ID(),
	})
	errorHandler(err)

	return api, nil
}
