package handler

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"helpers/awsservice"
	"helpers/dstruct"
	"helpers/dynamodbops"
	"helpers/models"
)

func Authenticate(credentials dstruct.UserLoginCredentials) error {
	// Get Salt
	svc := awsservice.GetDBConn()
	tableName := "UserAuth"
	keyAttributes := map[string]*dynamodb.AttributeValue{
		"Username": {
			S: aws.String(credentials.Username),
		},
	}

	entry := models.UserAuth{}
	result, err := dynamodbops.ReadItems(svc, tableName, keyAttributes)
	if err != nil {
		log.Error("Username does not exists: ", err)
		return fmt.Errorf("Username does not exists: %s", err)
	} else {
		err = dynamodbattribute.UnmarshalMap(result.Item, &entry)
		if err != nil {
			log.Error("Failed to unmarshal DB entry: %s", err)
		}
	}

	// Hash salt and pwd
	saltedPwd := entry.Salt + entry.HashedPwd
	fmt.Println(saltedPwd)

	// Check if hashedpwd is correct

	return nil
}
