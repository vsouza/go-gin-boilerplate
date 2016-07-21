# Go Gin Boilerplate
> A starter project with Golang, Gin and DynamoDB

[![Build Status][travis-image]][travis-url]
[![codebeat badge](https://codebeat.co/badges/ed248580-942c-4ffc-919f-d3681d28a799)](https://codebeat.co/projects/github-com-vsouza-go-gin-boilerplate)
[![Go Version][go-image]][go-url]
[![License][license-image]][license-url]


Golang Gin boilerplate with DynamoDB resource. Supports multiple configuration environments.

This project use a [DynamoDB base docker image](https://github.com/vsouza/docker-dynamoDB-local).

Setup DynamoDB dependency:  `docker run -d -p 8080:8080 vsouza/dynamo-local --port 8080`

![](header.jpg)

### Structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Installation

__Install Godeps__

`go get github.com/tools/godep`

then run:

```sh
make deps
```

## Usage example

`curl http://localhost:8888/health`

## Development setup

Run docker DynamoDB image or setup you `AWS_CREDENTIALS`

## Release History

* 0.0.1
    * Configuration by environment, Auth and Log middlewares, User entity.

## Meta

Vinicius Souza – [@iamvsouza](https://twitter.com/iamvsouza) – hi@vsouza.com

Distributed under the MIT license. See [License](https://vsouza.mit-license.org) for more information.

[https://github.com/vsouza](https://github.com/vsouza)

[go-image]: https://img.shields.io/badge/Go--version-1.6-blue.svg
[go-url]: https://golang.org/doc/go1.6
[travis-image]: https://travis-ci.org/vsouza/go-gin-boilerplate.svg?branch=master
[travis-url]: https://travis-ci.org/vsouza/go-gin-boilerplate
[license-image]: https://img.shields.io/badge/License-MIT-blue.svg
[license-url]: https://vsouza.mit-license.org
