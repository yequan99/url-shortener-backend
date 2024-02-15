package dynamodbops

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Scanning for items in a DynamoDB Table
func ScanItems(svc *dynamodb.DynamoDB, tableName string, nonPartitionKeyAttributeName string, nonPartitionKeyAttributeValue string) ([]map[string]*dynamodb.AttributeValue, error) {
	input := &dynamodb.ScanInput{
		TableName:        aws.String(tableName),
		FilterExpression: aws.String("#attr = :value"),
		ExpressionAttributeNames: map[string]*string{
			"#attr": aws.String(nonPartitionKeyAttributeName),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":value": {
				S: aws.String(nonPartitionKeyAttributeValue),
			},
		},
	}
	result, err := svc.Scan(input)
	if err != nil {
		return nil, fmt.Errorf("Unable to find: %w", err)
	}
	return result.Items, nil
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

	// Check if the item is not found
	if result.Item == nil {
		return nil, fmt.Errorf("item not found in table %s", tableName)
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

	return nil
}

// // Getting Auto-Increment values from DynamoDB Table
// func GetAutoIncrementValue(svc *dynamodb.DynamoDB, tableName string, autoIncrementKey uint64) (uint64, error) {
// 	result, err := svc.UpdateItem(&dynamodb.UpdateItemInput{
// 		TableName: aws.String(tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"ID": {
// 				S: aws.String(autoIncrementKey),
// 			},
// 		},
// 		UpdateExpression: aws.String("SET auto_increment = auto_increment + :inc"),
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":inc": {
// 				N: aws.String("1"),
// 			},
// 		},
// 		ReturnValues: aws.String("UPDATED_NEW"),
// 	})
// 	if err != nil {
// 		return 0, fmt.Errorf("error updating auto-increment value: %s", err)
// 	}

// 	autoIncrementValue, err := strconv.ParseUint(*result.Attributes["auto_increment"].N, 10, 64)
// 	if err != nil {
// 		return 0, fmt.Errorf("error converting auto-increment value to int: %s", err)
// 	}

// 	return autoIncrementValue, nil
// }

// Deleting items in DynamoDB Table
func DeleteItems(svc *dynamodb.DynamoDB, tableName string, keyAttributes map[string]*dynamodb.AttributeValue) error {

	// Check if item exists
	_, err := ReadItems(svc, tableName, keyAttributes)
	if err != nil {
		return fmt.Errorf("Item cannot be deleted as item is not in DB")
	}

	input := &dynamodb.DeleteItemInput{
		Key:       keyAttributes,
		TableName: aws.String(tableName),
	}
	_, err = svc.DeleteItem(input)
	return err
}

// Update items in DynamoDB Table
func UpdateItems(svc *dynamodb.DynamoDB, tableName string, keyAttributes map[string]*dynamodb.AttributeValue, expressionAttributes map[string]*dynamodb.AttributeValue, changeAttributes []string) error {

	// Check if attribute exists in DB
	_, err := ReadItems(svc, tableName, keyAttributes)
	if err != nil {
		return err
	}

	// Creating UpdateExpression
	var updateExpression *string
	if len(changeAttributes) > 0 {
		for _, att := range changeAttributes {
			update := att + " = :" + att
			if updateExpression != nil {
				// Concatenate with AND if conditionExpression is already set
				*updateExpression += ", " + att + " = :" + att
			} else {
				// Initialize conditionExpression if it's nil
				update = "set " + update
				updateExpression = &update
			}
		}
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: expressionAttributes,
		TableName:                 aws.String(tableName),
		Key:                       keyAttributes,
		ReturnValues:              aws.String("UPDATED_NEW"),
		UpdateExpression:          updateExpression,
	}

	_, err = svc.UpdateItem(input)
	if err != nil {
		return fmt.Errorf("Item cannot be updated: ", err)
	}

	return nil
}
