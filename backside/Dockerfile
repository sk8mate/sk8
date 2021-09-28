# Production

FROM golang:alpine3.14

WORKDIR /home/app

COPY . .

RUN go build -o backside

ENTRYPOINT ./backside
