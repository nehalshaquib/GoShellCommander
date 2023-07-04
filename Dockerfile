# Stage 1: Build the Go application
FROM golang:alpine as builder

# Stage 1: Build the Go application
RUN apk update && apk add --no-cache git

# Set the working directory inside the container
WORKDIR $GOPATH/src/goshellcommander/

# Copy the entire project to the working directory
COPY . .

# Download the module dependencies
RUN go mod download

# Build the Go application and output the binary to /go/bin/goshellcommander
RUN go build -o /go/bin/goshellcommander .

# Stage 2: Create a minimal runtime image
FROM alpine

# Copy the binary from the builder stage to the runtime image
COPY  --from=builder /go/bin/goshellcommander /go/bin/goshellcommander

# Set the entrypoint for the container to execute the goshellcommander binary
ENTRYPOINT [ "/go/bin/goshellcommander" ]
