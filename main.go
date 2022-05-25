package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rhmdnd/compserv-poc/compserv"
	pb "github.com/rhmdnd/compserv-poc/compserv"
)

func main() {
	// Get a connection to the database so we can pass it in as a
	// dependency to the service implementations.
	c := parseConfig()
	connStr := getDatabaseConnectionString(c)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Established database connection. %v", db)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Printf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAssessmentResultServer(grpcServer, compserv.NewServer(db))
	log.Printf("gRPC server listening on %s", lis.Addr())
	grpcServer.Serve(lis)
}

func parseConfig() map[string]string {
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.name", "compliance")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// FIXME(rhmdnd): Expand support for other configuration paths (e.g.,
	// /etc/compserv/config.yaml)
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	db_host := viper.GetString("database.host")
	db_port := viper.GetString("database.port")
	db_username := viper.GetString("database.username")
	db_name := viper.GetString("database.name")
	db_secret_arn := viper.GetString("database.secret_arn")
	db_secret_region := viper.GetString("database.secret_region")

	if len(db_host) == 0 {
		log.Fatal("Database host not provided.")
		os.Exit(1)
	} else if len(db_secret_arn) == 0 {
		log.Fatal("Database password not provided as an ARN.")
		os.Exit(1)
	} else if len(db_username) == 0 {
		log.Fatal("Database username not provided.")
		os.Exit(1)
	} else if len(db_secret_region) == 0 {
		log.Fatal("Database secret region not provided.")
		os.Exit(1)
	}

	m := make(map[string]string)
	m["db_host"] = db_host
	m["db_port"] = db_port
	m["db_secret_arn"] = db_secret_arn
	m["db_username"] = db_username
	m["db_name"] = db_name
	m["secret_region"] = db_secret_region

	log.Printf("Loaded configuration file: %s", viper.ConfigFileUsed())
	return m

}

func getDatabaseConnectionString(v map[string]string) string {
	secret := getSecret(v["db_secret_arn"], v["secret_region"])
	connectionString := fmt.Sprintf("host=%s user=%s password=%s port=%s", v["db_host"], v["db_username"], secret, v["db_port"])
	return connectionString
}

func getSecret(secretName string, region string) string {
	var secretString, decodedBinarySecret string

	//Create a Secrets Manager client
	sess, err := session.NewSession()
	if err != nil {
		// Handle session creation error
		fmt.Println(err.Error())
		return secretString
	}
	svc := secretsmanager.New(sess,
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return secretString
	}

	// Decrypts secret using the associated KMS key.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	if result.SecretString != nil {
		secretString = *result.SecretString
		return secretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return secretString
		}
		decodedBinarySecret = string(decodedBinarySecretBytes[:len])
		return decodedBinarySecret
	}
}
