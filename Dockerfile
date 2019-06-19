FROM golang:1.8

WORKDIR /go/src/app
COPY . /go/src/app

# Need these for more complex Golang apps
#RUN go get -d -v ./...
#RUN go install -v ./...
RUN go get -u github.com/gorilla/mux

CMD ["go", "run", "main.go"]

EXPOSE 8000