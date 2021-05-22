FROM golang:alpine

LABEL name="groundKim" 

ENV GIN_MODE=release \
    GO111MODULE=on \
    PORT=8888

ENV TZ=Asia/Seoul
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY . /app
WORKDIR /app

RUN apk update
RUN go mod download
VOLUME /app/data /home/groundkim/church_youth/data

RUN go build main.go

EXPOSE 8888

ENTRYPOINT [ "./main" ]
