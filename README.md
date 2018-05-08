# go-examples

## Examples for Go-Lang. 
 
 * Hello World 
 * greet service using go-kit 

## How To
* cd ~/go-examples/src/service
* go build greeting_service.go
* go run greeting_service.go 

## Test:
* curl -XPOST -d'{"s":"Rohit "}' localhost:8081/greet
