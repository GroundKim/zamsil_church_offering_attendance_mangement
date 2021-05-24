FROM golang:alpine

LABEL name="groundKim" 

ENV GIN_MODE=release \
    GO111MODULE=on \
    PORT=8888


COPY . /app
WORKDIR /app

RUN apk update
RUN go mod download
VOLUME /app/data /home/groundkim/church_youth/data

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
    echo "Asia/Seoul" > /etc/timezone
    
RUN go build main.go

EXPOSE 8888

ENTRYPOINT [ "./main" ]
