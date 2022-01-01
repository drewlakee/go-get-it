# go-get-it 

Simple database query script for UNIX-terminal usage

- Supports MongoDB

## Quick start

Commands:

```shell
Usage of ggi:
  -c string
    	MongoDB collection name : string (default "none")
  -f string
    	MongoDB find json-filter : string (default "none")
  -i int
    	MongoDB document '_id' field : int64
  -is string
    	MongoDB document '_id' field : string (default "none")
  -l int
    	Limit for find command : int64
```

Usage examples: 

```shell
~/$ ggi -c myCollectionWithIntId -i 10001 | jq
{
  "_id": 10001
}

~/$ ggi -c myCollection -is desiredId | jq
{
  "_id": "desiredId"
}

~/$ ggi -c myCollection -f '{}' | jq
[
 {
   "_id": "desiredId"
 },
 {
   "_id": "otherId"
 }
]

~/$ ggi -c myCollection -f '{}' -l 1 | jq
[
 {
   "_id": "desiredId"
 }
]
```

## How to run script

1. Create config file ```$HOME/.ggi/config.json```

Example,
```json
{
  "host" : "host",
  "port" : 27017,
  "database" : "database",
  "username" : "username",
  "password" : "password",
  "caFile" : "/usr/local/share/ca-certificates/my-cert/CAs.pem"
} 
```

Here ```caFile``` is optional, may be used for TLS communication

2. ```go build ~/cmd/ggi/ggi.go && ./ggi``` or ```go run ~/cmd/ggi/ggi.go```