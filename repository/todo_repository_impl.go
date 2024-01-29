package repository

import (
	"context"
	"main/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type todoRepositoryImpl struct {
	Client    *dynamodb.Client
	TableName string
}

func NewTodoRepository(client *dynamodb.Client, table_name string) TodoRepository {
	return &todoRepositoryImpl{
		Client:    client,
		TableName: table_name,
	}
}

func (repos *todoRepositoryImpl) FindAll() (*[]entity.TodoEntity, error) {
	var todo []entity.TodoEntity

	response, err := repos.Client.Scan(context.Background(), &dynamodb.ScanInput{
		TableName: aws.String(repos.TableName),
	})

	if err != nil {
		return nil, err
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}
