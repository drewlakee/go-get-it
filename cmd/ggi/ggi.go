package main

import (
	"context"
	"flag"
	"fmt"
	"go-get-it/pkg/cli"
	"go-get-it/pkg/cli/response"
	"go-get-it/pkg/config"
	mongo3 "go-get-it/pkg/mongo"
	"go-get-it/pkg/processor"
	mongo2 "go-get-it/pkg/processor/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	collectionNameCommand := cli.NewCommand(
		"c",
		"none",
		"MongoDB collection name : string")
	documentIdInt64Command := cli.NewCommand(
		"i",
		int64(0),
		"MongoDB document '_id' field : int64")
	documentIdStringCommand := cli.NewCommand(
		"is",
		"none",
		"MongoDB document '_id' field : string")
	filterCommand := cli.NewCommand(
		"f",
		"none",
		"MongoDB find json-filter : string")
	limitCommand := cli.NewCommand(
		"l",
		int64(0),
		"Limit for find command : int64")

	collectionNameCommand.InputValue = flag.String(
		collectionNameCommand.Flag,
		collectionNameCommand.DefaultValue.(string),
		collectionNameCommand.Description)
	documentIdInt64Command.InputValue = flag.Int64(
		documentIdInt64Command.Flag,
		documentIdInt64Command.DefaultValue.(int64),
		documentIdInt64Command.Description)
	documentIdStringCommand.InputValue = flag.String(
		documentIdStringCommand.Flag,
		documentIdStringCommand.DefaultValue.(string),
		documentIdStringCommand.Description)
	filterCommand.InputValue = flag.String(
		filterCommand.Flag,
		filterCommand.DefaultValue.(string),
		filterCommand.Description)
	limitCommand.InputValue = flag.Int64(
		limitCommand.Flag,
		limitCommand.DefaultValue.(int64),
		limitCommand.Description)

	flag.Parse()

	if *(collectionNameCommand.InputValue.(*string)) == collectionNameCommand.DefaultValue {
		fmt.Println("Collection name missed. Use -c")
		return
	}

	if *(documentIdInt64Command.InputValue.(*int64)) == documentIdInt64Command.DefaultValue &&
		*(documentIdStringCommand.InputValue.(*string)) == documentIdStringCommand.DefaultValue &&
		*(filterCommand.InputValue.(*string)) == filterCommand.DefaultValue {
		fmt.Println("Some commands missed. Use --help")
		return
	}

	cfg := config.LoadScriptConfig()
	url := mongo3.BuildConnectionUrl(cfg)
	if url == "exit" {
		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(response.GenerateJsonResponse("result", err.Error()))
		return
	}

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Println(response.GenerateJsonResponse("result", err.Error()))
			return
		}
	}(client, context.TODO())

	collection := client.Database(cfg.Database).Collection(*(collectionNameCommand.InputValue.(*string)))

	var commandProcessor processor.CommandJsonProcessor

	if commandProcessor == nil {
		if (*(documentIdInt64Command).InputValue.(*int64)) != (*documentIdInt64Command).DefaultValue.(int64) {
			id := *(documentIdInt64Command).InputValue.(*int64)
			commandProcessor = mongo2.NewFindOneByIdInt64MongoCommandProcessor(collection, id)
		}
	}

	if commandProcessor == nil {
		if (*(documentIdStringCommand).InputValue.(*string)) != documentIdStringCommand.DefaultValue.(string) {
			id := *(documentIdStringCommand).InputValue.(*string)
			commandProcessor = mongo2.NewFindOneByIdStringMongoCommandProcessor(collection, id)
		}
	}

	if commandProcessor == nil {
		if *(filterCommand.InputValue.(*string)) != filterCommand.DefaultValue.(string) {
			filter := *(filterCommand.InputValue.(*string))
			limit := *(limitCommand.InputValue.(*int64))
			commandProcessor = mongo2.NewFindFilterMongoCommandProcessor(collection, filter, limit)
		}
	}

	if commandProcessor == nil {
		fmt.Println("Incorrect command processing. Use --help")
	} else {
		fmt.Println(commandProcessor.GetJsonProcessResult())
	}
}
