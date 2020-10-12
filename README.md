Accounts microservice

This project is a little example how to structure a microservice. This is an attempt to use clean patterns applied in order to have a well tested microservice. 

## How to read this

### CLI
In the "cli" folder are the commands that can be used with this service.

### internal
The main code resides in this folder.

### QA
The QA tests will be inside this folder.

### tools
This is where are the tools needed for compile and test our code.

## How to run
```
# installing tools to compile
cd tools
go run mage.go

# building
make build

# running
# this will expose a graphQL server on 0.0.0.0:8080 (can be tested on the browser)
./bin/app serve

# running (using go)
go run cli/main.go serve

# testing using curl
# this performs a query
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"# Write your query or mutation here\nquery {\n  accounts {\n    id\n    holder\n  }\n}"}' --compressed

# this performs a mutation
curl 'http://localhost:8080/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H 'Origin: http://localhost:8080' --data-binary '{"query":"mutation {\n  createAccount(input: {\n    bankName: \"HSBC\",\n    accountNumber: \"123456789012345678\",\n    holder: \"CRL\"\n  }) {\n    id\n  }\n}"}' --compressed

# Have fun testing
# By default this uses a memorydb but a postgresql could be used too.

```


## TODO
* [x] Create a model
* [x] Create db interfaces
* [x] Create custom errors
* [x] Create cases (bussines rules)
* [x] Add graphql server
* [x] Add postgresql handler
* [x] Add memorydb 
* [x] Add migrations tool
* [x] Improve documentation
* [ ] Test graphql marshall and unmarshall
* [ ] Add QA test cases
* [ ] Add REST API dispatcher
* [ ] Improve tools installation process
* [ ] Improve magefile and makefile
* [ ] Add integration tests for postgresql db
* [ ] Improve Database interface
* [ ] Add cache interface
* [ ] Add GRPC comunication between microservices