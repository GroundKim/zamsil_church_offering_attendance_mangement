FROM golang:alpine

LABEL name="groundKim" 

ENV GIN_MODE=release \
    GO111MODULE=on \
    PORT=8888

COPY . /app
WORKDIR /app

RUN apk update
RUN go mod download

RUN go build main.go

EXPOSE 8888

ENTRYPOINT [ "./main" ]
