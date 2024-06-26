# Local Storage Server

## Overview
This project is a local storage server built with Go, providing file upload and download functionality over the network. The server also serves a simple Vue.js frontend for interacting with the API.

## Features
- Upload files to the server.
- Download files from the server.
- Simple health check endpoint.
- Vue.js frontend served over the `/` path.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```