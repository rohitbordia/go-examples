# go-examples

Examples for Go-Lang. 
1. Hello World 
2. greet service using go-kit 

cd ~/go-examples/src/service
go build greeting_service.go
go run greeting_service.go 

Test:
curl -XPOST -d'{"s":"Rohit "}' localhost:8081/greet
