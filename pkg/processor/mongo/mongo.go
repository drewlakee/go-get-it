package mongo

import (
	"context"
	"encoding/json"
	"go-get-it/pkg/cli/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneByIdInt64MongoCommandProcessor struct {
	collection *mongo.Collection
	id         int64
}

type FindOneByIdStringMongoCommandProcessor struct {
	collection *mongo.Collection
	id         string
}

type FindFilterMongoCommandProcessor struct {
	collection *mongo.Collection
	filter     string
	limit      int64
}

func NewSingleDocumentFetchByIdInt64MongoCommandProcessor(collection *mongo.Collection, id int64) *FindOneByIdInt64MongoCommandProcessor {
	return &FindOneByIdInt64MongoCommandProcessor{collection, id}
}

func NewSingleDocumentFetchByIdStringMongoCommandProcessor(collection *mongo.Collection, id string) *FindOneByIdStringMongoCommandProcessor {
	return &FindOneByIdStringMongoCommandProcessor{collection, id}
}

func NewFilterMongoCommandProcessor(collection *mongo.Collection, filter string, limit int64) *FindFilterMongoCommandProcessor {
	return &FindFilterMongoCommandProcessor{collection, filter, limit}
}

func (processor *FindOneByIdInt64MongoCommandProcessor) GetJsonProcessResult() string {
	var findOneResult bson.M
	err := processor.collection.FindOne(context.TODO(), bson.M{"_id": processor.id}).Decode(&findOneResult)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	parsedResult, err := json.Marshal(findOneResult)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	return string(parsedResult)
}

func (processor *FindOneByIdStringMongoCommandProcessor) GetJsonProcessResult() string {
	var fetchResult bson.M
	err := processor.collection.FindOne(context.TODO(), bson.M{"_id": processor.id}).Decode(&fetchResult)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	parsedResult, err := json.Marshal(fetchResult)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	return string(parsedResult)
}

func (processor *FindFilterMongoCommandProcessor) GetJsonProcessResult() string {
	var bsonFilter interface{}
	err := bson.UnmarshalExtJSON([]byte(processor.filter), false, &bsonFilter)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	findOptions := options.FindOptions{
		Limit: &processor.limit,
	}

	cursor, err := processor.collection.Find(context.TODO(), bsonFilter, &findOptions)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	var filterResult []bson.M
	err = cursor.All(context.TODO(), &filterResult)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	parsedResult, err := json.Marshal(filterResult)
	if err != nil {
		return response.GenerateJsonResponse("result", err.Error())
	}

	return string(parsedResult)
}
