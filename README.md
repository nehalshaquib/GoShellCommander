# GoShellCommander
GoShellCommander is a Go REST API that allows you to execute shell commands and retrieve their output. It provides a simple and convenient way to interact with the underlying operating system through HTTP requests.

## Features

- Execute shell commands and retrieve their output
- RESTful API endpoints for command execution
- Dockerized deployment using Docker Compose

## Prerequisites

- Go 1.19 or higher
- Docker (optional, for running the application in a container)

## Getting Started

1. Clone the repository:

	```shell
	git clone https://github.com/nehalshaquib/GoShellCommander.git
	```

2.  Change into the project directory:

	```shell
	cd GoShellCommander
	```
3.  Create a  .env  file in the project directory and set the following variables:
	```
	TOKENS=token1,token2,token3,...
	PORT=8085
	GIN_MODE=debug
	```

4.  Build the Go application:
	```shell
	make build
	```

5.  Run the application:
	```shell
	make run
	```

6.  The GoShellCommander API will be available at  http://localhost8085.

## API Endpoints

  **Endpoint:** `/api/cmd`

**Method:** POST

**Description:** Execute a shell command and retrieve its output.

**Header**: `token`: `<your_token_value>`

**Request Body**
```json
{
    "command_name": "docker",
    "arguments": ["ps", "-a"]
}
```

|Field  | Type | Description |
|--|--|--|
|  `command_name` | `string` | The name of the command to be executed |
|  `arguments` | `array`| Optional arguments for the command |

**Response:**

In case of a successful execution, the response body will be:

```json
{
    "result": "command_output"
}
```
In case of an error, the response body will be:
```json
{
    "detail": "error detail",
    "error": "error"
}
```

## Configuration

The application can be configured using environment variables. Create a .env file in the project directory and set the following variables:

-   PORT: The port on which the server will listen (default: 8085)
-   GIN_MODE: The mode in which the Gin framework will run (default: debug)
-   TOKENS: A comma-separated list of authorized tokens for authentication

## Testing

run the tests for the Go application, use the following command:

```shell
make test
```
## Docker Deployment

To deploy the application using Docker Compose, use the following command:

```shell
make up
```
This will start the application in a Docker container.

**Note**:  Please note that when running the GoShellCommander API server within a Docker container, the executed commands will be performed within the Docker shell environment rather than the host shell. If you prefer to execute commands within the host shell, please utilize the `make build` and `make run` commands to build and run the Go binaries directly on the host machine.
