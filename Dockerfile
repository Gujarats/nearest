FROM golang:latest

# Set the application directory
COPY . /go/src/github.com/Gujarats/API-Golang
WORKDIR /go/src/github.com/Gujarats/API-Golang

RUN go get -v ./
RUN go build

# Copy our code from the current folder to /app inside the container
ADD . /app

# Make port 80 available for links and/or publish
EXPOSE 8080
