# Go HelloWorld API

This is a simple "Hello World" API written in Go. It serves as a basic example of how to create a HTTP server that responds with a JSON message.

## Features

- Simple HTTP server
- JSON response
- Single endpoint that returns a "Hello World" message

## Installation

1. Clone the repository:
```git clone https://github.com/yamamoto99/go-helloworld-api.git```

2. Navigate to the project directory:
`cd go-helloworld-api`

## Usage

1. Run the server:
go run main.go

2. The server will start running on `http://localhost:8080`

3. Test the API using curl:
curl http://localhost:8080

You should receive the following response:
```json
{"message": "Hello World"}
```
API Endpoint

- GET /: Returns a JSON object with a "Hello World" message
