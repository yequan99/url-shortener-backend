package dynamodbops

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	UserID string `json:"userID"`
	Hash   string `json:"hash"`
}

// Listing items in a DynamoDB Table
func ReadItems(svc *dynamodb.DynamoDB, tableName string, keyAttributes map[string]*dynamodb.AttributeValue) (*dynamodb.GetItemOutput, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       keyAttributes,
	})
	if err != nil {
		return nil, fmt.Errorf("error calling GetItem: %s", err)
	}

	return result, nil
}

// Inserting items into DynamoDB Table
func InsertItems(svc *dynamodb.DynamoDB, tableName string, newItem interface{}, conditions []string) error {
	attribute, err := dynamodbattribute.MarshalMap(newItem)
	if err != nil {
		return fmt.Errorf("error marshalling new item: %s", err)
	}

	// Creating conditional expression
	var conditionExpression *string
	if len(conditions) > 0 {
		for _, condition := range conditions {
			con := "attribute_not_exists(" + condition + ")"
			if conditionExpression != nil {
				// Concatenate with AND if conditionExpression is already set
				*conditionExpression += " AND " + con
			} else {
				// Initialize conditionExpression if it's nil
				conditionExpression = &con
			}
		}
	}

	// Insert only if condition does not exist
	input := &dynamodb.PutItemInput{
		Item:                attribute,
		TableName:           aws.String(tableName),
		ConditionExpression: conditionExpression,
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return fmt.Errorf("error calling PutItem: %s", err)
	}

	fmt.Printf("Successfully added item to table %s\n", tableName)
	return nil
}
