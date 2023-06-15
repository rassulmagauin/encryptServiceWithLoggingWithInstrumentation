# Encrypt Service with Logging and Instrumentation

This repository contains a service written in Go that includes logging and instrumentation functionalities.

## Project Structure

The project primarily consists of Go (.go) files, along with a module (.mod) and a checksum (.sum) file. Here's a brief overview of the main files:

- `main.go`: This is the entry point of the application. It sets up the logging, metrics, and the service handlers.

- `helpers/endpoins.go`: Contains the definitions for various endpoints of the service.

- `helpers/implementations.go`: Contains the implementation of various functions or methods used in the service.

- `helpers/instrumentation.go`: Contains code related to the instrumentation of the service, such as performance monitoring or logging.

- `helpers/jsonutils.go`: Contains utility functions for working with JSON data.

- `helpers/middleware.go`: Contains middleware functions, which are used to handle requests and responses in the service.

- `helpers/models.go`: Contains the data models used in the service.

## API Endpoints

The service has two main endpoints:

- `/encrypt`: This endpoint accepts a key and a text as input and returns the encrypted text. If there's an error during the process, it returns the error message.

- `/decrypt`: This endpoint accepts a key and an encrypted message as input and returns the decrypted text. If there's an error during the process, it returns the error message.

In addition, the service also provides a `/metrics` endpoint for Prometheus metrics.

## Getting Started

To run this project, you will need to have Go installed on your machine. Once you have Go installed, you can clone this repository and run the `main.go` file.

```bash
git clone https://github.com/rassulmagauin/encryptServiceWithLoggingWithInstrumentation.git
cd encryptServiceWithLoggingWithInstrumentation
go run main.go
```
