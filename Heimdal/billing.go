package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createBillingAlarm(ctx *pulumi.Context) error {

	topic, err := sns.NewTopic(ctx, "ragnarAlert", &sns.TopicArgs{
		Name: pulumi.String("ragnarALERT"),
	})
	errorHandler(err)

	_, err = sns.NewTopicSubscription(ctx, "ragnarTopicSub", &sns.TopicSubscriptionArgs{
		Topic:    topic.Arn,
		Protocol: pulumi.String("email"),
		Endpoint: pulumi.String("foxhound@the0x.dev"),
	})
	errorHandler(err)

	_, err = cloudwatch.NewMetricAlarm(ctx, "ragnarBillin", &cloudwatch.MetricAlarmArgs{
		Name:               pulumi.String("ragnar-billing"),
		ComparisonOperator: pulumi.String("GreaterThanOrEqualToThreshold"),
		EvaluationPeriods:  pulumi.Int(1),
		MetricName:         pulumi.String("EstCharges"),
		Namespace:          pulumi.String("AWS/Billing"),
		Period:             pulumi.Int(21600),
		Statistic:          pulumi.String("Maximum"),
		Threshold:          pulumi.Float64(10),
		AlarmDescription:   pulumi.String("alerts when price is too damn high"),
		ActionsEnabled:     pulumi.Bool(true),
		AlarmActions: pulumi.Array{
			topic.Arn,
		},
		Dimensions: pulumi.StringMap{
			"Currency": pulumi.String("USD"),
		},
	})
	errorHandler(err)

	return nil
}
