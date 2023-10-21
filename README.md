# Learning hexagonal architecture in Go #1

source: https://www.golinuxcloud.com/hexagonal-architectural-golang/

# How to Use

- run `go run cmd/main.go` or `go run cmd/main.go --db redis`
- from terminal run 
    - post message: `curl -X POST http://localhost:5000/messages -H 'Content-Type: application/json' -d '{"body": "message from curl"}'` 
    - get all messages: `curl -X GET http://localhost:5000/messages -H 'Content-Type: application/json'`