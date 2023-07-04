FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/goshellcommander/
COPY . .
RUN go mod download
RUN go build -o /go/bin/goshellcommander .

FROM alpine
COPY  --from=builder /go/bin/goshellcommander /go/bin/goshellcommander
ENTRYPOINT [ "/go/bin/goshellcommander" ]

