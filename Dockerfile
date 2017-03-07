FROM golang:latest


COPY . /go/src/github.com/Gujarats/API-Golang 
WORKDIR /go/src/github.com/Gujarats/API-Golang 

RUN go get ./
RUN go build

# if dev setting will use pilu/fresh for code reloading via docker-compose volume sharing with local machine
# if production setting will build binary
CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi


EXPOSE 8080
