## base image of peoject
FROM golang:1.17.1-alpine3.14
## directory where our app going to live
RUN mkdir /app
RUN apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
## add all project files
COPY ./ /app
## compile the binery and create executable file
RUN go build -o chatroom .
## Our start command which kicks off
## our newly created binary executable
## Expose port 8080 to the outside world
EXPOSE 8080
CMD ["/app/chatroom"]
