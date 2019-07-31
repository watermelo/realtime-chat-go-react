FROM golang:1.12 AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build .
CMD ["/app/realtime-chat-go-react"]
