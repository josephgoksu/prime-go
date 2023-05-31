# prime-go

[![Build and Push an Image to ECR](https://github.com/josephgoksu/prime-go/actions/workflows/main.yml/badge.svg)](https://github.com/josephgoksu/prime-go/actions/workflows/main.yml)

[![CodeQL](https://github.com/josephgoksu/prime-go/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/josephgoksu/prime-go/actions/workflows/github-code-scanning/codeql)

This is a simple repository to test the Go language. It will be used for platform engineering exercises.

It will be used in [this](https://github.com/josephgoksu/platform-eng-exercise/) repository:

## Table of Contents

- [prime-go](#prime-go)
  - [Table of Contents](#table-of-contents)
  - [How to run](#how-to-run)
  - [How to test](#how-to-test)
  - [How to build](#how-to-build)
  - [How to run the binary](#how-to-run-the-binary)
  - [How to run docker](#how-to-run-docker)

## How to run

```bash
go run main.go
```

## How to test

```bash
go test
```

## How to build

```bash
go build
```

## How to run the binary

```bash
./prime-go
```

## How to run docker

```bash
docker build -t prime-go .
```

```bash
docker run -p 8080:8080 prime-go:latest
```

Author: [Joseph Goksu](https://github.com/josephgoksu)
