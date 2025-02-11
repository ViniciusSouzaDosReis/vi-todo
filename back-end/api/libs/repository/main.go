package repository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RepositoryHello struct {
	connection *dynamodb.Client
	tableName  string
}

type Hello struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const tableName = "Hello"

func NewRepositoryHello(connection *dynamodb.Client) *RepositoryHello {
	return &RepositoryHello{
		connection,
		tableName,
	}
}

func (r *RepositoryHello) GetHello(ctx context.Context) (Hello, error) {
	result, err := r.connection.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{
				Value: "1",
			},
		},
	})

	if err != nil {
		return Hello{}, err
	}

	var entity Hello
	err = attributevalue.UnmarshalMap(result.Item, &entity)
	if err != nil {
		return Hello{}, err
	}

	return entity, nil
}
