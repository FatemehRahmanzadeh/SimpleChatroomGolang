## base image of peoject
FROM golang:1.17.1-alpine3.14
## directory where our app going to live
RUN mkdir /app
ENV CGO_ENABLED=0
## add all project files
COPY . /app
WORKDIR /app

RUN go mod download
## compile the binery and create executable file
RUN go build -o chatroom .
## Our start command which kicks off
## our newly created binary executable
## Expose port 8080 to the outside worldor
EXPOSE 8080
CMD ["/app/chatroom"]


