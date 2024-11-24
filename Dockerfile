FROM golang:1.23.3
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api-server
EXPOSE 8003
CMD ["/go-api-server"]