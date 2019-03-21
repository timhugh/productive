FROM golang:latest

WORKDIR /go/src/productive
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80/tcp
CMD ["productive"]