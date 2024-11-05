/*
Copyright © 2024 FOXHOUND0x foxhound@the0x.dev
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	username   string
	password   string
	userPoolID string
	clientID   string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Logins into AWS cognito to get a token via PKCE",
	Long: `Logins into AWS cognito to get a token via PKCE in order
	to authenticate with the AWS API Gateway.`,
	Run: func(cmd *cobra.Command, args []string) {
		if username == "" {
			fmt.Printf("Enter Username: ")
			fmt.Scanln(&username)
		}
		if password == "" {
			fmt.Printf("Enter Password: ")
			fmt.Scanln(&password)
			bytePassword, _ := term.ReadPassword(int(os.Stdin.Fd()))
			password = string(bytePassword)
		}

		_, err := cognitoAuth(username, password)
		if err != nil {
			fmt.Printf("Authentication failed: %v\n:", err)
			return
		}

		fmt.Println("Successfully Authenticated")
	},
}

func cognitoAuth(username, password string) (string, error) {
	conf, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("failed to load default AWS configuraiton from ~/.aws/config", err)
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(conf)
	if userPoolID == "" || clientID == "" {
		return "", fmt.Errorf("Cognito User pool ID and Client ID must be set")
	}

	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: cognitoidentityprovider.AuthFlowTypeUserPasswordAuth,
		ClientId: &clientID,
		AuthParameters: map[string]string{
			"USERNAME": username,
			"PASSWORD": password,
		},
	}

	authOutput, err := cognitoClient.InitiateAuth(context.TODO(), authInput)
	if err != nil {
		return "", fmt.Errorf("auth failed: %w", err)
	}

	if authOutput.AuthenticationResult == nil || authOutput.AuthenticationResult.IdToken == nil {
		return "", fmt.Errorf("Token is Null bro")
	}

	return *authOutput.AuthenticationResult.IdToken, nil
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username for Cognito")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password for Cognito")

	userPoolID = os.Getenv("AWS_COGNITO_USER_POOL_ID")
	clientID = os.Getenv("AWS_COGNITO_CLIENT_ID")

	loginCmd.Flags().StringVar(&userPoolID, "user-pool-id", userPoolID, "AWS Cognito User Pool ID")
	loginCmd.Flags().StringVar(&clientID, "client-id", clientID, "AWS Cognito Client ID")
}
