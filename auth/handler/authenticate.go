package handler

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"golang.org/x/crypto/bcrypt"

	"helpers/awsservice"
	"helpers/dstruct"
	"helpers/dynamodbops"
	"helpers/models"
)

func Authenticate(credentials dstruct.UserLoginCredentials) error {
	// Get Hash
	svc := awsservice.GetDBConn()
	tableName := "UserAuth"
	keyAttributes := map[string]*dynamodb.AttributeValue{
		"Username": {
			S: aws.String(credentials.Username),
		},
	}

	entry := models.UserAuth{}
	hashedPwd, err := dynamodbops.ReadItems(svc, tableName, keyAttributes)
	if err != nil {
		log.Error("Username does not exists: ", err)
		return fmt.Errorf("Username does not exists: %s", err)
	} else {
		err = dynamodbattribute.UnmarshalMap(hashedPwd.Item, &entry)
		if err != nil {
			log.Error("Failed to unmarshal DB entry: %s", err)
		}
	}

	// Check if hashedpwd is correct
	authenticated := comparePasswords(entry.HashedPwd, credentials.Password)
	if !authenticated {
		return fmt.Errorf("[User Login] Failed authentication")
	}

	return nil
}

func Register(credentials dstruct.UserLoginCredentials) error {

	fmt.Println(credentials.Username)
	fmt.Println(credentials.Password)

	// Salt and hash password
	hash := hashAndSalt([]byte(credentials.Password))
	fmt.Println(hash)

	// Insert to DB
	svc := awsservice.GetDBConn()

	tableName := "UserAuth"
	item := models.UserAuth{
		Username:  credentials.Username,
		HashedPwd: hash,
	}
	condition := []string{"Username"}
	err := dynamodbops.InsertItems(svc, tableName, item, condition)
	if err != nil {
		log.Error("[User Registration] Unable to register user:", err)
		return fmt.Errorf("[User Registration] Unable to register user:", err)
	}

	return nil
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
