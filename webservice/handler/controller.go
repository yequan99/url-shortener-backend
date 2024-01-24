package handler

import (
	"fmt"

	"helpers/awsservice"
	"helpers/dynamodbops"

	"helpers/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

func GetShortURL(username string, longURL string) (string, error) {

	svc := awsservice.GetDBConn()
	domain := "http://127.0.0.1:8080/"

	// Check if LongURL already present
	tableName := "UrlTable"
	keyAttributes := map[string]*dynamodb.AttributeValue{
		"LongURL": {
			S: aws.String(longURL),
		},
	}

	var shortID string
	tableItem, err := dynamodbops.ReadItems(svc, tableName, keyAttributes)
	if tableItem != nil {
		// LongURL entry already present
		entry := models.UrlTable{}
		err = dynamodbattribute.UnmarshalMap(tableItem.Item, &entry)
		if err != nil {
			log.Errorf("Failed to unmarshal DB entry: %s", err)
			return "", fmt.Errorf("error when unmarshalling")
		} else {
			shortID = entry.ShortID
		}
	} else {
		// LongURL not present and need to generate short ID
		for {
			sid, err := shortid.New(1, shortid.DefaultABC, 2342)
			shortID, err = sid.Generate()
			if err != nil {
				log.Errorf("Error generating shortid")
				return "", err
			}

			// Check if Short ID already used
			tableName = "UrlTable"
			nonPartitionKeyAttributeName := "ShortID"
			nonPartitionKeyAttributeValue := shortID

			scanItems, err := dynamodbops.ScanItems(svc, tableName, nonPartitionKeyAttributeName, nonPartitionKeyAttributeValue)
			var foundItems []models.UrlTable
			err = dynamodbattribute.UnmarshalListOfMaps(scanItems, &foundItems)
			if err != nil {
				log.Errorf("Failed to unmarshal DynamoDB attributes:", err)
			}
			if len(foundItems) == 0 {
				// ShortID not used
				break
			}
		}
	}

	// Store data in UserURL Table
	tableName = "UserURL"
	userUrl_item := models.UserURL{
		Username: username,
		ShortURL: domain + shortID,
		LongURL:  longURL,
	}

	err = dynamodbops.InsertItems(svc, tableName, userUrl_item, []string{})
	if err != nil {
		log.Errorf("[Short URL Generation] Unable to create new URL: %s", err)
		return "", fmt.Errorf("Unable to create new URL")
	}

	// Store data in UrlTable Table
	tableName = "UrlTable"
	urlTable_item := models.UrlTable{
		LongURL:  longURL,
		ShortID:  shortID,
		ShortURL: domain + shortID,
	}

	err = dynamodbops.InsertItems(svc, tableName, urlTable_item, []string{})
	if err != nil {
		log.Errorf("[Short URL Generation] Unable to create new URL: %s", err)
		return "", fmt.Errorf("Unable to create new URL")
	}

	return domain + shortID, nil
}
