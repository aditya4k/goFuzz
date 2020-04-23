[![Build Status][build-badge]][build-url] [![Go Report Card][goreport-card-badge]][goreport-card] [![Go Coverage][codecov-card-badge]][codecov-card]

# goFuzz
This is an client server solution to find out the payloads that can be used to bypass
Web Application Firewall. Client can use the cli to make the request to server to retrieve
the payload corresponding to an attack type (xss|sql) against a particular site protected
by an WAF.

Inspiration Project: [*WAFNinja*](https://github.com/khalilbijjou/WAFNinja)

## Go Dependency Setup
Go to the root directory of the source code and then run the following command if
go dependency is not set up for the project: 
`export $GOPATH && export GO114MODULE=on && go mod init && go get ./...`


## High Level Design

## Client Build


## Server Deployment


## CLI


## Contributing

